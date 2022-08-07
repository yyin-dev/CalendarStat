package event_stat

import (
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/yinfredyue/CalendarStat/authenticate"
	"github.com/yinfredyue/CalendarStat/cmd/flags"
	"github.com/yinfredyue/CalendarStat/stat"
)

type config struct {
	startDate    time.Time
	endDate      time.Time
	credential   string
	groupEventBy stat.GroupEventBy
}

func Flags() []cli.Flag {
	startDate := flags.DateFlag("start-date")
	endDate := flags.DateFlag("end-date")
	credentialFlag := flags.CredentialFlag()
	groupEventByFlag := flags.GroupEventByFlag()

	return []cli.Flag{
		&startDate,
		&endDate,
		&credentialFlag,
		&groupEventByFlag,
	}
}

func configFromCliContext(ctx *cli.Context) *config {
	// Check groupEventBy
	allowedSet := make(map[string]stat.GroupEventBy)
	allowedSet["calendar"] = stat.ByCalendar
	allowedSet["colorId"] = stat.ByColorId

	groupEventByFlag := ctx.String("group-event-by")
	groupEventBy, ok := allowedSet[groupEventByFlag]
	if !ok {
		log.Fatal("Invalid group-event-by flag")
		return nil
	}

	return &config{
		startDate:    *ctx.Timestamp("start-date"),
		endDate:      *ctx.Timestamp("end-date"),
		credential:   ctx.String("credential"),
		groupEventBy: groupEventBy,
	}
}

func Cmd(ctx *cli.Context) error {
	conf := configFromCliContext(ctx)
	service := authenticate.GetService(conf.credential)

	calendarsResult, err := service.CalendarList.List().Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, calendar := range calendarsResult.Items {
		eventsRes, err := service.Events.
			List(calendar.Id).
			ShowDeleted(false).
			SingleEvents(true).
			TimeMin(conf.startDate.Format(time.RFC3339)).
			TimeMax(conf.endDate.Format(time.RFC3339)).
			OrderBy("startTime").
			MaxResults(2500). // TODO: This limit is unlikely to be hit. But it's possible. Maybe we need paging?
			Do()
		if err != nil {
			log.Fatal(err)
			return err
		}

		events := eventsRes.Items
		eventGroup := stat.BuildEventGroup(events, calendar, conf.groupEventBy, conf.startDate, conf.endDate)
		eventGroupStat := eventGroup.Stat()

		fmt.Printf("%v\n  %v\n\n", eventGroup, eventGroupStat)
	}

	return nil
}
