package main

import (
	"github.com/urfave/cli/v2"
	"github.com/yinfredyue/CalendarStat/cmd/calendars"
	"github.com/yinfredyue/CalendarStat/cmd/colors"
	"log"
	"os"
)

type runConfig struct {
	credential string
}

func run(config runConfig) {
	//service := authenticate.GetService(config.credential)
	//
	//// Get events
	//startTimeString := config.startDate.Format(time.RFC3339)
	//endTimeString := config.endDate.Format(time.RFC3339)
	//
	//eventsRes, err := service.Events.
	//	List("yyin5@andrew.cmu.edu").
	//	ShowDeleted(false).
	//	SingleEvents(true).
	//	TimeMin(startTimeString).
	//	TimeMax(endTimeString).
	//	OrderBy("startTime").
	//	Do()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//events := eventsRes.Items
	//
	//analyze.Analyze(events)
}

func main() {

	app := &cli.App{
		Name:  "CalendarStat",
		Usage: "The tool for analyzing your Google Calendar",
		//Flags: []cli.Flag{
		//	&cli.StringFlag{
		//		Name:  "start-date",
		//		Value: time.Now().AddDate(0, 0, -7).Format(time.RFC3339),
		//		Usage: "Start date of the range you want to analyze",
		//	},
		//	&cli.StringFlag{
		//		Name:  "end-date",
		//		Value: time.Now().Format(time.RFC3339),
		//		Usage: "Start date of the range you want to analyze",
		//	},
		//	credentialFlag,
		//},
		Commands: []*cli.Command{
			{
				Name:  "colors",
				Usage: "show all colorIds",
				Flags: colors.Flags(),
				Action: func(ctx *cli.Context) error {
					return colors.Cmd(ctx)
				},
			},
			{
				Name:  "calendars",
				Usage: "show all calendars",
				Flags: calendars.Flags(),
				Action: func(ctx *cli.Context) error {
					return calendars.Cmd(ctx)
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			//startDateStr := ctx.String("start-date")
			//endDateStr := ctx.String("end-date")
			//
			//startDate, err := time.Parse(time.RFC3339, startDateStr)
			//if err != nil {
			//	return err
			//}
			//
			//endDate, err := time.Parse(time.RFC3339, endDateStr)
			//if err != nil {
			//	return err
			//}
			//
			//config := runConfig{
			//	startDate:  startDate,
			//	endDate:    endDate,
			//	credential: ctx.String("credential"),
			//}
			//
			//run(config)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
