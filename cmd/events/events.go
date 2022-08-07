package events

import (
	"github.com/urfave/cli/v2"
	"github.com/yinfredyue/CalendarStat/authenticate"
	"github.com/yinfredyue/CalendarStat/cmd/flags"
	"github.com/yinfredyue/CalendarStat/utils"
	"log"
	"time"
)

type config struct {
	startDate  time.Time
	endDate    time.Time
	credential string
	calendarId string
}

func Flags() []cli.Flag {
	startDate := flags.DateFlag("start-date")
	endDate := flags.DateFlag("end-date")
	credential := flags.CredentialFlag()
	calendarId := flags.CalendarIdFlag()

	return []cli.Flag{
		&credential,
		&startDate,
		&endDate,
		&calendarId,
	}
}

func configFromCliContext(ctx *cli.Context) *config {
	return &config{
		credential: ctx.String("credential"),
		startDate:  *ctx.Timestamp("start-date"),
		endDate:    *ctx.Timestamp("end-date"),
		calendarId: ctx.String("calendar-id"),
	}
}

func Cmd(ctx *cli.Context) error {
	conf := configFromCliContext(ctx)
	service := authenticate.GetService(conf.credential)

	eventsRes, err := service.Events.
		List(conf.calendarId).
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
	for _, event := range events {
		utils.PrintEvent(event)
	}
	return nil
}
