package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `himma : a lightweight container engine used in Edge computing gateway.`

func main() {
	app := cli.NewApp()
	app.Name = "himma"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
		listCommand,
		logCommand,
		execCommand,
		stopCommand,
		removeCommand,
		commitCommand,
		networkCommand,
	}

	app.Before = func(context *cli.Context) error {

		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
