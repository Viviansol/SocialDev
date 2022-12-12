package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecuteTemplatesA(w http.ResponseWriter, template string, datas interface{}) {
	templates.ExecuteTemplate(w, template, datas)
}
