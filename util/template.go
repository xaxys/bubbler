package util

import (
	"bytes"
	"text/template"
)

func ExecuteTemplate(templ, templateName string, fns template.FuncMap, data interface{}) string {
	tmpl := template.Must(template.New("").Funcs(fns).Parse(templ))
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, templateName, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
