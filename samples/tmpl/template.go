package template

import "strings"

//go:generate togo tmpl -func funcMap -format html

var funcMap = map[string]interface{}{
	"uppercase": strings.ToUpper,
	"lowercase": strings.ToLower,
}
