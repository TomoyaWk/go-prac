package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var paramDesc string
	var paramAmount int

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
			{
				Name:  "add",
				Usage: "add new expense.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "description",
						Value:       "no description",
						Usage:       "description for expense",
						Destination: &paramDesc,
					},
					&cli.IntFlag{
						Name:        "amount",
						Value:       0,
						Usage:       "expense amount",
						Destination: &paramAmount,
					},
				},
				Action: func(ctx *cli.Context) error {
					//test
					fmt.Println("description:", paramDesc)
					fmt.Println("amount:", paramAmount)

					created, err := CreateExpense(paramDesc, paramAmount)
					if err != nil {
						fmt.Printf("failed!: %v", err)
						return nil
					}
					fmt.Printf("Expense added successfully (ID: %d)", created.ID)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
