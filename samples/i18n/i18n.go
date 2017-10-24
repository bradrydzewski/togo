package i18n

//go:generate goi18n -flat -outdir files files/en-us.yaml
//go:generate goi18n -flat -outdir files files/fr-fr.yaml files/en-us.yaml
//go:generate goi18n -flat -outdir files files/es-es.yaml files/en-us.yaml
//go:generate goi18n -flat -outdir files files/zh-cn.yaml files/en-us.yaml

//go:generate goi18n -flat -outdir files files/fr-fr.all.json files/fr-fr.untranslated.json files/en-us.all.json
//go:generate goi18n -flat -outdir files files/fr-fr.all.json files/es-es.untranslated.json files/en-us.all.json
//go:generate goi18n -flat -outdir files files/fr-fr.all.json files/zh-cn.untranslated.json files/en-us.all.json

//go:generate togo i18n
