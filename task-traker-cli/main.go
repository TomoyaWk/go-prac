package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
					if desc == "" {
						println("please enter task description.")
						return nil
					}

					newTask, err := createNewTask(desc)
					if err != nil {
						fmt.Printf("failed: %q", err)
						return err
					}
					fmt.Printf("Task added successfully (ID: %d)", newTask.Id)
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "update task description.",
				Action: func(ctx *cli.Context) error {
					argId := ctx.Args().Get(0)
					desc := ctx.Args().Get(1)
					if desc == "" {
						println("please enter task description.")
						return nil
					}
					id, err := strconv.Atoi(argId)
					if err != nil {
						fmt.Printf("invalid argment: %s is not interger.", argId)
						return nil
					}

					updated, err := UpdateTask(id, desc)
					if err != nil {
						fmt.Printf("failed: %q", err)
						return err
					}
					fmt.Printf("Task updated successfully: (ID: %d, desc: %s)", updated.Id, updated.Description)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
