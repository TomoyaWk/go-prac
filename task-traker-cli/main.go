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
					query := ctx.Args().First()
					GetTasks(query)
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
					id, err := strconv.Atoi(argId) //intパース
					if err != nil {
						fmt.Printf("invalid argment: %s is not interger.", argId)
						return nil
					}

					updated, err := UpdateTaskDesc(id, desc)
					if err != nil {
						fmt.Printf("failed: %q", err)
						return err
					}
					fmt.Printf("Task updated successfully: (ID: %d, desc: %s)", updated.Id, updated.Description)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete task",
				Action: func(ctx *cli.Context) error {
					argId := ctx.Args().Get(0)

					id, err := strconv.Atoi(argId) //intパース
					if err != nil {
						fmt.Printf("invalid argment: %s is not interger.", argId)
						return nil
					}

					deleted, err := DeleteTask(id)
					if err != nil {
						fmt.Printf("failed: %q", err)
						return err
					}
					fmt.Printf("Task deleted successfully: (ID: %d, desc: %s)", deleted.Id, deleted.Description)
					return nil
				},
			},
			{
				Name:  "mark-in-progress",
				Usage: "update task status in-progress",
				Action: func(ctx *cli.Context) error {
					argId := ctx.Args().Get(0)
					id, err := strconv.Atoi(argId) //intパース
					if err != nil {
						fmt.Printf("invalid argment: %s is not interger.", argId)
						return nil
					}
					updated, err := UpdateTaskStatus(id, InProgress)
					if err != nil {
						fmt.Printf("failed: %q", err)
						return err
					}
					fmt.Printf("status updated successfully: (ID: %d, desc: %s)", updated.Id, updated.Description)
					return nil
				},
			},
			{
				Name:  "mark-done",
				Usage: "update task status Done",
				Action: func(ctx *cli.Context) error {
					argId := ctx.Args().Get(0)
					id, err := strconv.Atoi(argId) //intパース
					if err != nil {
						fmt.Printf("invalid argment: %s is not interger.", argId)
						return nil
					}
					updated, err := UpdateTaskStatus(id, Done)
					if err != nil {
						fmt.Printf("failed: %q", err)
						return err
					}
					fmt.Printf("status updated successfully: (ID: %d, desc: %s)", updated.Id, updated.Description)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
