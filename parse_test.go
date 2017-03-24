package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	stmts, err := parse(strings.NewReader(example))
	if err != nil {
		t.Error(err)
	}

	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	enc.Encode(stmts)
}

var example = `

--
-- +statement select-pets
--

select *
from pets
{{ if .Type }}
WHERE type = ?
{{ else if .MaxPrice }}
WHERE price <= ?
{{ end }}

-- +statement select-pets-by-id

select *
from pets
where id = ?

-- +statement select-pets-by-type postgres

select *
from pets
where type = $1

`
