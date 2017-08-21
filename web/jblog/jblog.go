package jblog

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/jekabolt/j1site/configurator"
	"github.com/jekabolt/j1site/jdblayer"
	"github.com/jekabolt/j1site/web/restapi"
)

var hydraWebTemplate *template.Template
var historylog struct {
	logs []string
	sync.RWMutex
}

func Run() error {
	var err error

	conf := struct {
		Filespath string   `json:"filespath"`
		Templates []string `json:"templates"`
	}{}
	err = configurator.GetConfiguration(configurator.JSON, &conf, "./web/portalconfig.json")
	if err != nil {
		return err
	}

	hydraWebTemplate, err = template.ParseFiles(conf.Templates...)
	if err != nil {
		return err
	}

	restapi.InitializeAPIHandlers()
	log.Println(conf.Filespath)
	log.Println(conf.Templates)
	fs := http.FileServer(http.Dir(conf.Filespath))
	http.Handle("/", fs)
	http.HandleFunc("/blog/", blogHandler)
	http.HandleFunc("/404/", notFoundHandler)
	http.HandleFunc("/datahon/", blogDatahon)
	http.ListenAndServe(":8080", nil)
	fmt.Println("kek")
	return err
}

func crewhandler(w http.ResponseWriter, r *http.Request) {
	dblayer, err := jdblayer.ConnectDatabase("mongodb", "localhost")
	if err != nil {
		return
	}
	all, err := dblayer.AllMembers()
	if err != nil {
		return
	}
	err = hydraWebTemplate.ExecuteTemplate(w, "crew.html", all)
	if err != nil {
		log.Println(err)
	}
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	about := struct {
		Msg string `json:"message"`
	}{}
	err := configurator.GetConfiguration(configurator.JSON, &about, "./hydraweb/about.json")
	if err != nil {
		return
	}
	err = hydraWebTemplate.ExecuteTemplate(w, "about.html", about)
	if err != nil {
		log.Println(err)
	}
}
func blogHandler(w http.ResponseWriter, r *http.Request) {
	err := hydraWebTemplate.ExecuteTemplate(w, "blog.html", nil)
	if err != nil {
		log.Println(err)
	}
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := hydraWebTemplate.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		log.Println(err)
	}
}
func blogDatahon(w http.ResponseWriter, r *http.Request) {
	err := hydraWebTemplate.ExecuteTemplate(w, "datahon.html", nil)
	if err != nil {
		log.Println(err)
	}
}
