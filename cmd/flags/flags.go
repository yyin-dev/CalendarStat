package flags

import (
	"github.com/urfave/cli/v2"
	"time"
)

func CredentialFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "credential",
		Value: "../credentials.json",
		Usage: "Google OAuth credential file",
	}
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
