package main

import (
	"dummyserver/src/internal/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:  false,
		FullTimestamp:  true,
		DisableSorting: true,
	})
	log.Info("Starting dummy server...")
	app.StartServer()
}
