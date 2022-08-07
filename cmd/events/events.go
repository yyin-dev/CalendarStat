package events

import (
	"github.com/urfave/cli/v2"
	"github.com/yinfredyue/CalendarStat/analyze"
	"github.com/yinfredyue/CalendarStat/authenticate"
	"github.com/yinfredyue/CalendarStat/cmd/flags"
	"log"
	"time"
)

type config struct {
	startDate  time.Time
	endDate    time.Time
	credential string
}

func Flags() []cli.Flag {
	credential := flags.CredentialFlag()
	startDate := flags.DateFlag("start-date")
	endDate := flags.DateFlag("end-date")

	return []cli.Flag{
		&credential,
		&startDate,
		&endDate,
	}
}

func configFromCliContext(ctx *cli.Context) *config {
	return &config{
		credential: ctx.String("credential"),
		startDate:  *ctx.Timestamp("start-date"),
		endDate:    *ctx.Timestamp("end-date"),
	}
}

func Cmd(ctx *cli.Context) error {
	conf := configFromCliContext(ctx)
	service := authenticate.GetService(conf.credential)

	eventsRes, err := service.Events.
		List("yyin5@andrew.cmu.edu").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(conf.startDate.Format(time.RFC3339)).
		TimeMax(conf.endDate.Format(time.RFC3339)).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Fatal(err)
	}
	events := eventsRes.Items

	analyze.Analyze(events)
	return nil
}
