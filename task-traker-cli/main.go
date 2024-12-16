package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "task-cli",
		Usage: "task tracker",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "task list",
				Action: func(ctx *cli.Context) error {
					GetTasks()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
