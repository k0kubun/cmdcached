package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandServer,
	commandClient,
}

var commandServer = cli.Command{
	Name:  "server",
	Usage: "",
	Description: `
`,
	Action: doServer,
}

var commandClient = cli.Command{
	Name:  "client",
	Usage: "",
	Description: `
`,
	Action: doClient,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doServer(c *cli.Context) {
}

func doClient(c *cli.Context) {
}
