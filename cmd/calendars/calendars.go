package calendars

import (
	"github.com/urfave/cli/v2"
	"github.com/yinfredyue/CalendarStat/authenticate"
	"github.com/yinfredyue/CalendarStat/cmd/flags"
	"github.com/yinfredyue/CalendarStat/utils"
	"log"
)

type config struct {
	credential string
}

func Flags() []cli.Flag {
	credentialFlag := flags.CredentialFlag()

	return []cli.Flag{
		&credentialFlag,
	}
}

func configFromCliContext(ctx *cli.Context) *config {
	return &config{
		credential: ctx.String("credential"),
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
		utils.PrintCalendarListEntry(calendar)
	}

	return nil
}
