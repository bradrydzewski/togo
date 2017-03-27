package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sqlbin"
	app.Usage = "sqlbin provides too for embedding sql opterations in your binary"
	app.Version = "1.0.0"
	app.Author = "bradrydzewski"
	app.Commands = []cli.Command{
		ddlCommand,
		sqlCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
