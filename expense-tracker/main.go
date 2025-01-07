package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "expense-tracker",
		Usage: "expense tracker",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "expense list",
				Action: func(c *cli.Context) error {
					GetExpenses()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
