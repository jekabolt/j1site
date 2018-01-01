package jblog

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

var tpl *template.Template

func SetHandlers(r *chi.Mux, fp string, tpls []string) {

	var err error
	tpl, err = template.ParseFiles(tpls...)
	if err != nil {
		log.Printf("[ERROR] template.ParseFiles %s \n", err.Error())
	}

	filesDir := http.Dir(fp)
	FileServer(r, "/static", filesDir)

	r.Get("/", index)
	r.Get("/404", notFonudHandler)
	r.Get("/blog", blog)
	r.Get("/datahon", datahon)
	// TemplatesFinder(r, ".", "templates")
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("[ERROR] tpls.ExecuteTemplate %s \n", err.Error())
	}
}
func notFonudHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		log.Printf("[ERROR] tpls.ExecuteTemplate %s \n", err.Error())
	}
}
func blog(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "blog.html", nil)
	if err != nil {
		log.Printf("[ERROR] tpls.ExecuteTemplate %s \n", err.Error())
	}
}

func datahon(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "datahon.html", nil)
	if err != nil {
		log.Printf("[ERROR] tpls.ExecuteTemplate %s \n", err.Error())
	}
}
