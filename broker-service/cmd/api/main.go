package main

import (
	"broker/logger"
	"fmt"
	"log"

	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {

	// init the logger
	logger.InitLogger("/var/log/broker.log")

	app := Config{}

	log.Printf("starting broker service on port %s", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		logger.Log.Panic(err)
		// log.Panic(err)
	}
}
