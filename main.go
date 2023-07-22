package main

import (
	"log"
	"mercury-x/cmd"
	"os"

	"github.com/urfave/cli/v2"
)

const ExitCodeInvalidUsage = 2

func build() *cli.App {
	app := cli.NewApp()
	app.Usage = "채용 플랫폼 크롤러"
	app.Suggest = true

	// When showing default help, exit with an error code
	app.Action = func(c *cli.Context) error {
		var err error

		args := c.Args()
		if args.Present() {
			err = cli.ShowCommandHelp(c, args.First())
		} else {
			err = cli.ShowAppHelp(c)
		}

		if err != nil {
			return err
		}
		return cli.Exit("", ExitCodeInvalidUsage)
	}

	app.Commands = cmd.GetCommands()

	return app
}

func run(args ...string) error {
	return build().Run(args)
}

func main() {
	err := run(os.Args...)
	if err != nil {
		if msg := err.Error(); msg != "" {
			log.Println(msg)
		}
		if v, ok := err.(cli.ExitCoder); ok {
			os.Exit(v.ExitCode())
		}
		os.Exit(ExitCodeInvalidUsage)
	}
}
