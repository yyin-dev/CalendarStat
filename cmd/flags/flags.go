package flags

import (
	"time"

	"github.com/urfave/cli/v2"
)

func CredentialFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "credential",
		Value: "../secret/credentials.json",
		Usage: "Google OAuth credential file",
	}
}

func CalendarIdFlag() cli.StringFlag {
	return cli.StringFlag{Name: "calendar-id",
		Usage:    "Calendar ID (view all calendar IDs with command 'calendars'",
		Required: true}
}

func GroupEventByFlag() cli.StringFlag {
	return cli.StringFlag{Name: "group-event-by",
		Usage:    "Choose from [calendar|colorId]",
		Required: true}
}

func DateFlag(name string) cli.TimestampFlag {
	return cli.TimestampFlag{
		Name:        name,
		DefaultText: "",
		Required:    true,
		Value:       nil,
		Timezone:    nil,
		Layout:      time.RFC3339,
	}
}
