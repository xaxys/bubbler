package util

import (
	"bytes"
	"text/template"
)

func ExecuteTemplate(templ, templateName string, fns template.FuncMap, data interface{}) string {
	if fns == nil {
		fns = make(template.FuncMap)
	}
	fns["panic"] = templatePanic
	fns["calc"] = templateCalc
	fns["dict"] = templateDict
	fns["iterate"] = templateIterate
	fns["ToPascalCase"] = ToPascalCase
	fns["TocamelCase"] = TocamelCase
	fns["Tosnake_case"] = Tosnake_case
	fns["ToALLCAP_CASE"] = ToALLCAP_CASE
	tmpl := template.Must(template.New("").Funcs(fns).Parse(templ))
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, templateName, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func templatePanic(v interface{}) error {
	panic(v)
}

func templateIterate(startValue, endValue any) []int64 {
	start := ToInt64(startValue)
	end := ToInt64(endValue)
	var r []int64
	for i := start; i < end; i++ {
		r = append(r, i)
	}
	return r
}

func templateDict(keysAndValues ...interface{}) map[string]interface{} {
	if len(keysAndValues)%2 != 0 {
		panic("dict must have even number of arguments")
	}
	d := make(map[string]interface{})
	for i := 0; i < len(keysAndValues); i += 2 {
		k, ok := keysAndValues[i].(string)
		if !ok {
			panic("dict keys must be strings")
		}
		d[k] = keysAndValues[i+1]
	}
	return d
}

func templateCalc(args ...interface{}) interface{} {
	return Calculate(args...)
}
