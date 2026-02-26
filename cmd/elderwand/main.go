package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	routerpkg "github.com/kyotuca/elderwand/cmd/http"
)

func render(writer http.ResponseWriter, name string, data any) {
	tpl := template.Must(template.ParseFiles(
		"views/layout/base.html",
		"views/partials/nav.html",
		"views/pages/"+name+".html",
	))

	if err := tpl.ExecuteTemplate(writer, "base", data); err != nil {
		http.Error(writer, err.Error(), 500)
	}
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

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app.css", http.FileServer(http.Dir("public")))
	app := routerpkg.New(render)
	mux.Handle("/", app)
	log.Println("http://localhost:8081")
	_ = http.ListenAndServe("localhost:8081", mux)
}
