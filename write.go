package main

import (
	"io"
	"strings"
	"text/template"
)

var funcs = map[string]interface{}{
	"join": func(a, b string) string {
		if b == "" {
			return a
		}
		return strings.Join([]string{a, b}, "-")
	},
	"camelize": camelize,
}

type data struct {
	Package    string
	Labels     []string
	Statements []*statement
}

func generate(w io.Writer, pkg string, stmts []*statement) error {
	t, err := template.New("_").Funcs(funcs).Parse(text)
	if err != nil {
		panic(err)
	}
	set := map[string]string{}
	for _, stmt := range stmts {
		if stmt.Driver != "" {
			set[stmt.Driver] = stmt.Driver
		}
	}
	var labels []string
	for k := range set {
		labels = append(labels, k)
	}
	return t.Execute(w, data{
		Package:    pkg,
		Labels:     labels,
		Statements: stmts,
	})
}

var text = `
package {{ .Package }}

// Lookup returns the named statement.
func Lookup(name string) (string) {
   return index[name]
}

{{ if .Labels }}
// LookupTag returns the named statement by tag.
func LookupTag(name, tag string) (string) {
  switch tag {
  {{- range .Labels }}
  case {{ printf "%q" . }}:
    return {{ . }}Index[name]
  {{- end }}
  default:
   return index[name]
  }
}
{{ end }}

var index = map[string]string{
  {{- range $i, $stmt := .Statements -}}
    {{- if not .Driver }}
      {{- $name := join .Name .Driver }}
      {{ printf "%q" $stmt.Name }}: {{ camelize $name }},
    {{- end -}}
  {{ end }}
}

{{ $statements := .Statements }}
{{ range $ii, $label := .Labels }}
var {{ $label }}Index = map[string]string{
  {{- range $i, $stmt := $statements -}}
    {{- if eq $label $stmt.Driver }}
      {{- $name := join .Name .Driver }}
      {{ printf "%q" $stmt.Name }}: {{ camelize $name }},
    {{ end -}}
  {{- end -}}
}
{{ end }}

{{ range .Statements -}}
{{ $name := join .Name .Driver }}
var {{ camelize $name }} = ` + "`" + "\n{{.Value }}\n" + "`" + `

{{ end }}
`

func camelize(kebab string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}

// old version of template

var textBackup = `
package {{ .Package }}

import "text/template"

type statement struct {
  name  string
  value string
  templ *template.Template
}

func init() {
  for _, s := range statements {
    s.templ = template.Must(
      template.New(s.name).Parse(s.value),
    )
  }
}

// Lookup returns the named statement.
func Lookup(name string) (s string) {
   stmt, ok := index[name]
   if ok {
     return stmt.value
   }
   return
}

// LookupTag returns the named statement by tag.
func LookupTag(name, tag string) (s string) {
  switch tag {
  {{- range .Labels }}
  case {{ printf "%q" . }}:
    stmt, ok := {{ . }}Index[name]
    if ok {
      return stmt.value
    }
  {{- end }}
  default:
    // no-op
  }
  // fallback to default statement
  return Lookup(name)
}

var statements = []*statement{
  {{ range .Statements -}}
    {{ $name := join .Name .Driver -}}
    {
      name:  {{ printf "%q" .Name }},
      value: {{ camelize $name }},
    },
  {{ end }}
}

var index = map[string]*statement{
  {{- range $i, $stmt := .Statements -}}
    {{- if not .Driver }}
      {{ printf "%q" $stmt.Name }}: statements[{{$i}}],
    {{- end -}}
  {{ end }}
}

{{ $statements := .Statements }}
{{ range $ii, $label := .Labels }}
var {{ $label }}Index = map[string]*statement{
  {{- range $i, $stmt := $statements -}}
    {{- if eq $label $stmt.Driver }}
      {{ printf "%q" $stmt.Name }}: statements[{{$i}}],
    {{ end -}}
  {{- end -}}
}
{{ end }}

{{ range .Statements -}}
{{ $name := join .Name .Driver }}
var {{ camelize $name }} = ` + "`" + "\n{{.Value }}\n" + "`" + `

{{ end }}
`
