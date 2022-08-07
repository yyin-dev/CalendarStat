package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yinfredyue/CalendarStat/cmd/calendars"
	"github.com/yinfredyue/CalendarStat/cmd/colors"
	"github.com/yinfredyue/CalendarStat/cmd/events"
)

func main() {
	colorCommand := &cli.Command{
		Name:  "colors",
		Usage: "Show all colorIds",
		Flags: colors.Flags(),
		Action: func(ctx *cli.Context) error {
			return colors.Cmd(ctx)
		},
	}

	calendarsCommand := &cli.Command{
		Name:  "calendars",
		Usage: "Show all calendars",
		Flags: calendars.Flags(),
		Action: func(ctx *cli.Context) error {
			return calendars.Cmd(ctx)
		},
	}

	eventsCommand := &cli.Command{
		Name:  "events",
		Usage: "Show all events",
		Flags: events.Flags(),
		Action: func(ctx *cli.Context) error {
			return events.Cmd(ctx)
		},
	}

	app := &cli.App{
		Name:  "CalendarStat",
		Usage: "A tool for analyzing your Google Calendar",
		Commands: []*cli.Command{
			colorCommand,
			calendarsCommand,
			eventsCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
