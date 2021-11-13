package main

import (
	_ "flag"
	"log"
	"os"

	"github.com/urfave/cli"
)

// var name = flag.String("name", "stranger", "your wonderful name")
// var age = flag.Int("age", 18, "your current age")

func main() {
	// flag.Parse()
	// log.Printf("Hello %s (%d years), Welcome to the command line world", *name, *age)

	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name:        "name",
			Usage:       "your real name",
			Value:       "stranger",
		},
		cli.IntFlag{
			Name:        "age",
			Usage:       "your real age",
			Value:       18,
		},
	}

	app.Action = func(c *cli.Context) error {
		log.Printf("Hello %s (%d years), Welcome to the command line world", c.String("name"), c.Int("age"))
		return nil
	}

	app.Run(os.Args)
}