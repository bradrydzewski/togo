package main

import "github.com/urfave/cli"

var embedCommand = cli.Command{
	Name:  "embed",
	Usage: "embed commands",
	Subcommands: []cli.Command{
		ddlCommand,
		sqlCommand,
	},
}
