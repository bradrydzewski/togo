package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sqlbin"
	app.Usage = "sqlbin provides command line tools for generating embedded sql opterations"
	app.Version = "1.0.0-alpha"
	app.Commands = []cli.Command{
		ddlCommand,
		sqlCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
