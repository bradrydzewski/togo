package i18n

import "github.com/nicksnyder/go-i18n/i18n"

// list of embedded template files.
var files = []struct {
	name string
	data string
}{
	{
		name: "en-us.all.json",
		data: enUsall,
	}, {
		name: "es-es.all.json",
		data: esEsall,
	}, {
		name: "fr-fr.all.json",
		data: frFrall,
	}, {
		name: "zh-cn.all.json",
		data: zhCnall,
	},
}

func init() {
	for _, file := range files {
		i18n.ParseTranslationFileBytes(file.name, []byte(file.data))
	}
}

//
// embedded template files.
//

// files/en-us.all.json
var enUsall = `{
  "greeting": {
    "other": "Hello"
  },
  "welcome": {
    "other": "Welcome"
  }
}`

// files/es-es.all.json
var esEsall = `{
  "greeting": {
    "other": ""
  },
  "welcome": {
    "other": ""
  }
}`

// files/fr-fr.all.json
var frFrall = `{
  "greeting": {
    "other": "Bonjour"
  },
  "welcome": {
    "other": "Bienvenue"
  }
}`

// files/zh-cn.all.json
var zhCnall = `{
  "greeting": {
    "other": "Hello"
  },
  "welcome": {
    "other": "Welcome"
  }
}`
