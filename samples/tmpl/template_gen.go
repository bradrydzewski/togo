package template

import "html/template"

// list of embedded template files.
var files = []struct {
	name string
	data string
}{
	{
		name: "en.tmpl",
		data: en,
	}, {
		name: "es.tmpl",
		data: es,
	},
}

// T exposes the embedded templates.
var T *template.Template

func init() {
	T = template.New("_").Funcs(funcMap)
	for _, file := range files {
		T = template.Must(
			T.New(file.name).Parse(file.data),
		)
	}
}

//
// embedded template files.
//

// files/en.tmpl

var en = `<html>
<body>
Hello world
</body>
</html>
`

// files/es.tmpl

var es = `<html>
<body>
Hola Mundo
</body>
</html>
`
