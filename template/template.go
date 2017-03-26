package template

//go:generate go-bindata -pkg template -o template_gen.go files/

import (
	"bytes"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

// parsed template
var tmpl *template.Template

// init loads the templates from the embedded file map. This function will not
// compile if go generate is not executed before.
func init() {
	dir, _ := AssetDir("files")
	tmpl = template.New("_").Funcs(FuncMap)
	for _, name := range dir {
		path := filepath.Join("files", name)
		src := MustAsset(path)
		tmpl = template.Must(
			tmpl.New(name).Parse(string(src)),
		)
	}
}

// Execute renders the named template and writes to io.Writer wr.
func Execute(wr io.Writer, name string, data interface{}) error {
	buf := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		return nil
	}
	src, err := format(buf)
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, src)
	return err
}

// FuncMap provides extra functions for the templates.
var FuncMap = template.FuncMap{
	"substr":   substr,
	"camelize": camelize,
}

func substr(s string, i int) string {
	return s[:i]
}

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
