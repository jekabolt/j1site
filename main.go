package main

import (
	"fmt"
	"os"

	"github.com/jekabolt/j1site/jlogger"
	"github.com/jekabolt/j1site/web/jblog"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no port")
		os.Exit(1)
	}
	port := os.Args[1]

	logger := jlogger.GetInstance()
	logger.Println("Starting web service")
	err := jblog.Run(port)
	if err != nil {
		logger.Println("Could not run web portal", err)
	}
}
