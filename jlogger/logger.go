package jlogger

import (
	"log"
	"os"
	"sync"
)

type jLogger struct {
	*log.Logger
	filename string
}

var jlogger *jLogger
var once sync.Once

//GetInstance create a singleton instance of the logger
func GetInstance() *jLogger {
	once.Do(func() {
		jlogger = createLogger("logger.log")
	})
	return jlogger
}

//Create a logger instance
func createLogger(fname string) *jLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &jLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra ", log.Lshortfile),
	}
}
