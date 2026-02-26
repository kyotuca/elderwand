package main

import (
	"html/template"
	"net/http"
	"strings"
	"time"
)

type Page struct {
	Title  string
	Body   []byte
	Player string
}

type Router struct {
	render func(http.ResponseWriter, string, any)
}

func loadPage() *Page {
	body := []byte{0}
	return &Page{Title: "foo", Body: body, Player: "bar"}
}

func processTemplate() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"upper": strings.ToUpper,
		"truncate": func(s string, n int) string {
			if len(s) <= n {
				return s
			}
			return s[:n] + "..."
		},
		"fmtDate": func(t time.Time) string {
			return t.Format("01 Jan 1970")
		},
	}).ParseGlob("views/**/*.html"))
}
