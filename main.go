package main

import (
	"os"
	"github.com/codegangsta/cli"
)

const (
	bundleID = "redfit.alfred-google-translate-workflow"
	version  = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "alfred-translation"
	app.Version = version
	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Action: cmdSetup,
		},
		{
			Name:   "translate",
			Action: cmdTranslate,
		},
	}

	loadConfig()

	app.Run(os.Args)
}
