package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	prefix  = "-- +statement "
	comment = "--"
	newline = "\n"
)

type statement struct {
	Name   string
	Value  string
	Driver string
}

func parse(r io.Reader) ([]*statement, error) {
	var (
		stmts []*statement
		stmt  *statement
	)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, prefix) {
			stmt = new(statement)
			stmt.Name, stmt.Driver = parsePrefix(line)
			stmts = append(stmts, stmt)
		}
		if strings.HasPrefix(line, comment) {
			continue
		}
		if stmt != nil {
			stmt.Value += line + newline
		}
	}
	for _, stmt := range stmts {
		stmt.Value = strings.TrimSpace(stmt.Value)
	}
	return stmts, nil
}

func parsePrefix(line string) (name string, driver string) {
	line = strings.TrimPrefix(line, prefix)
	line = strings.TrimSpace(line)
	fmt.Sscanln(line, &name, &driver)
	return
}
