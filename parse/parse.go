package parse

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	prefix    = "-- +statement "
	comment   = "--"
	newline   = "\n"
	delimiter = ";"
)

// Statement represents a statment in the sql file.
type Statement struct {
	Name   string
	Value  string
	Driver string
}

// // Parser parses the sql file.
// type Parser struct {
// 	prefix string
// }
//
// // New returns a new parser with the given prefix.
// func New(prefix string) *Parser {
// 	return &Parser{prefix: prefix}
// }

// Parse parses the sql file and returns a list of statements.
func Parse(r io.Reader) ([]*Statement, error) {
	var (
		stmts []*Statement
		stmt  *Statement
	)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, prefix) {
			stmt = new(Statement)
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
