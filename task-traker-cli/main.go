package main

import (
	"fmt"
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
			{
				Name:  "add",
				Usage: "create New task",
				Action: func(ctx *cli.Context) error {
					desc := ctx.Args().First()
					newTask, err := createNewTask(desc)
					if err != nil {
						fmt.Printf("failed: %q", err)
						return err
					}
					fmt.Printf("new Task created: %s", newTask.Description)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
