package main

import (
	"log"
	"os"
	router2 "paydemo/router"
)

func main() {
	router := router2.Router()

	err := router.Run(":80")
	if err != nil {
		log.Println("start service failed")
		os.Exit(0)
	}
}
