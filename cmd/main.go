package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/yinfredyue/CalendarStat/analyze"
	"github.com/yinfredyue/CalendarStat/authenticate"
	"github.com/yinfredyue/CalendarStat/utils"
	"log"
)

type runConfig struct {
	startDate  time.Time
	endDate    time.Time
	credential string
}

func run(config runConfig) {
	service, colorService := authenticate.GetService(config.credential)

	// Get color configs
	colorConfig, err := colorService.Get().Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Color config contains %v colors for events\n", len(colorConfig.Event))

	// Get calendars
	calendarsResult, err := service.CalendarList.List().Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, calendar := range calendarsResult.Items {
		utils.PrintCalendarListEntry(calendar)
	}

	// Get events
	startTimeString := config.startDate.Format(time.RFC3339)
	endTimeString := config.endDate.Format(time.RFC3339)

	eventsRes, err := service.Events.
		List("yyin5@andrew.cmu.edu").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(startTimeString).
		TimeMax(endTimeString).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Fatal(err)
	}
	events := eventsRes.Items

	analyze.Analyze(events)
}

func main() {
	app := &cli.App{
		Name:  "CalendarStat",
		Usage: "The tool for analyzing your Google Calendar",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "start-date",
				Value: time.Now().AddDate(0, 0, -7).Format(time.RFC3339),
				Usage: "Start date of the range you want to analyze",
			},
			&cli.StringFlag{
				Name:  "end-date",
				Value: time.Now().Format(time.RFC3339),
				Usage: "Start date of the range you want to analyze",
			},
			&cli.StringFlag{
				Name:  "credential",
				Value: "./credentials.json",
				Usage: "Google OAuth credential file",
			},
		},
		Action: func(ctx *cli.Context) error {
			startDateStr := ctx.String("start-date")
			endDateStr := ctx.String("end-date")

			startDate, err := time.Parse(time.RFC3339, startDateStr)
			if err != nil {
				return err
			}

			endDate, err := time.Parse(time.RFC3339, endDateStr)
			if err != nil {
				return err
			}

			config := runConfig{
				startDate:  startDate,
				endDate:    endDate,
				credential: ctx.String("credential"),
			}

			run(config)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
