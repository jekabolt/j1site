package restapi

import (
	"log"
	"net/http"

	"github.com/jekabolt/j1site/configurator"
)

type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

func InitializeAPIHandlers() {
	conf := new(DBlayerconfig)
	err := configurator.GetConfiguration(configurator.JSON, conf, "./web/apiconfig.json")
	if err != nil {
		log.Fatal("Error decoding JSON", err)
	}
	h := newhydraCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		log.Fatal("Error connecting to db ", err)
	}
	http.HandleFunc("/hydracrew/", h.handleHydraCrewRequests)
}

func RunAPI() {
	InitializeAPIHandlers()
	http.ListenAndServe(":8061", nil)
}
