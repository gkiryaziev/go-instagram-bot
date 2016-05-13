package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"

	"github.com/gkiryaziev/go-instagram-bot/core"
)

// checkError check errors
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "IBot"
	app.Usage = "Instagram bot"
	app.Version = "0.1.3"

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run bot",
			Action: func(c *cli.Context) error {
				coreService, err := core.NewService()
				checkError(err)
				err = coreService.Run()
				checkError(err)
				return nil
			},
		},
		{
			Name:  "droptables",
			Usage: "Drop database tables",
			Action: func(c *cli.Context) error {
				coreService, err := core.NewService()
				checkError(err)
				err = coreService.DropTables()
				checkError(err)
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "Migrate database",
			Action: func(c *cli.Context) error {
				coreService, err := core.NewService()
				checkError(err)
				err = coreService.Migrate()
				checkError(err)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
