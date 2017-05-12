package main

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/bradrydzewski/togo/template"
)

var httpCommand = cli.Command{
	Name:   "http",
	Usage:  "generate an http filesystem",
	Action: tmplAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "package",
			Value: "http",
		},
		cli.StringFlag{
			Name:  "input",
			Value: "files/**",
		},
		cli.StringFlag{
			Name:  "output",
			Value: "http_gen.go",
		},
		cli.StringFlag{
			Name:  "trim-prefix",
			Value: "files",
		},
		cli.StringSliceFlag{
			Name:  "plain-text",
			Value: &cli.StringSlice{"html", "js", "css"},
		},
	},
}

func httpAction(c *cli.Context) error {
	pattern := c.Args().First()
	if pattern == "" {
		pattern = c.String("input")
	}

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _ = range matches {

	}

	wr := os.Stdout
	if output := c.String("output"); output != "" {
		wr, err = os.Create(output)
		if err != nil {
			return err
		}
		defer wr.Close()
	}

	return template.Execute(wr, "http.tmpl", nil)
}
