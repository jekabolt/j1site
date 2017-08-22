package main

import (
	"github.com/jekabolt/j1site/jlogger"
	"github.com/jekabolt/j1site/web/jblog"
)

func main() {

	logger := jlogger.GetInstance()
	logger.Println("Starting web service")
	err := jblog.Run()
	if err != nil {
		logger.Println("Could not run web portal", err)
	}
}
