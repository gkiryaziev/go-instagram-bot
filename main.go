package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"

	"github.com/gkiryaziev/go-instagram-bot/core"
)

// checkError check errors
func checkError(err error) {
	log.Fatal(err)
}

func main() {
	app := cli.NewApp()
	app.Name = "IBot"
	app.Usage = "Instagram news bot"
	app.Version = "0.1.3"

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run service",
			Action: func(c *cli.Context) {
				coreSvc, err := core.NewService()
				checkError(err)
				err = coreSvc.Run()
				checkError(err)
			},
		},
		{
			Name:  "droptables",
			Usage: "Drop database tables",
			Action: func(c *cli.Context) {
				coreSvc, err := core.NewService()
				checkError(err)
				err = coreSvc.DropTables()
				checkError(err)
			},
		},
		{
			Name:  "migrate",
			Usage: "Migrate database",
			Action: func(c *cli.Context) {
				coreSvc, err := core.NewService()
				checkError(err)
				err = coreSvc.Migrate()
				checkError(err)
			},
		},
	}
	app.Run(os.Args)
}
