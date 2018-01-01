package main

import (
	"fmt"
	"log"

	"github.com/jekabolt/jsite"
)

func main() {

	jsite.SetUpLogger()

	conf, err := jsite.LoadConfiguration("jsite.config")
	if err != nil {
		log.Printf("[ERROR] jsite.LoadConfiguration %s \n", err)
	}

	fmt.Println(conf)

	site, err := jsite.Init(&conf)
	if err != nil {
		log.Printf("[ERROR] jsite.Init %s \n", err.Error())
	}

	err = site.Run()
	if err != nil {
		log.Printf("[ERROR] site.Run() %s \n", err.Error())
	}

}
