package internal

import "html/template"

func GetCustomFunction() template.FuncMap {
	return template.FuncMap{
		"str2html": func(raw string) template.HTML {
			return template.HTML(raw)
		},
	}
}
