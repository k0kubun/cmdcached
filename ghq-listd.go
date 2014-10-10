package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghq-listd"
	app.Version = Version
	app.Usage = ""
	app.Author = "Takashi Kokubun"
	app.Email = "takashikkbn@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
