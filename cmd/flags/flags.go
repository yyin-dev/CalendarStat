package flags

import "github.com/urfave/cli/v2"

func CredentialFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "credential",
		Value: "../credentials.json",
		Usage: "Google OAuth credential file",
	}
}
