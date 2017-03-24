package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	input   = flag.String("input", "", "input file name; required")
	output  = flag.String("o", "", "output file name; required")
	pkgname = flag.String("pkg", "main", "output package name; required")
)

func main() {
	flag.Parse()

	files, err := filepath.Glob(*input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	for _, file := range files {
		out, ferr := ioutil.ReadFile(file)
		if ferr != nil {
			fmt.Println(ferr)
			os.Exit(1)
		}
		buf.Write(out)
	}

	// parses the sql statements from the file.
	stmts, err := parse(&buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf.Reset()
	err = generate(&buf, *pkgname, stmts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// formats the generated file using gofmt
	pretty, err := format(&buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// create output source for file. defaults to
	// stdout but may be file.
	var out = os.Stdout
	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		defer out.Close()
	}

	io.Copy(out, pretty)
}
