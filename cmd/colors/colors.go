package colors

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
	"github.com/yinfredyue/CalendarStat/authenticate"
	"github.com/yinfredyue/CalendarStat/cmd/flags"
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
	colorService := authenticate.GetColorService(conf.credential)

	colorConfig, err := colorService.Get().Do()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(colorConfig.Event)
	fmt.Printf("Color config contains %v colors for events\n", len(colorConfig.Event))

	return nil
}
