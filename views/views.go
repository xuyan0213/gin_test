package views

import (
	"embed"
	"html/template"
)

var (
	embedTmpl embed.FS
	funcMap   = template.FuncMap{}
	GoTpl     = template.Must(template.New("").Funcs(funcMap).ParseGlob("*.html"))
)
