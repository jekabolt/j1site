package jsite

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jekabolt/jsite/web/jblog"
)

type Site struct {
	config *Config
	route  *chi.Mux
}

// Init initializes Multy instance
func Init(conf *Config) (*Site, error) {
	site := &Site{
		config: conf,
	}
	if err := site.initRoutes(conf); err != nil {
		log.Printf("[ERROR] site.initRoutes %s \n", err.Error())
	}
	return site, nil
}

func (site *Site) initRoutes(conf *Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	jblog.SetHandlers(router, conf.Filespath, conf.Templates)
	site.route = router

	return nil
}

// Run runs service
func (site *Site) Run() error {
	log.Printf("[WARN] Run service")
	http.ListenAndServe(site.config.Host+site.config.Port, site.route)
	return nil
}
