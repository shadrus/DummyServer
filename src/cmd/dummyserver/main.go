package main

import (
	"dummyserver/src/internal/app"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	port := getEnv("port", "8080")
	logPath := getEnv("logPath", "")
	if logPath != "" {
		f, err := os.OpenFile(path.Join(logPath, "dummyserver.log"), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		mw := io.MultiWriter(os.Stderr, f)
		log.SetOutput(mw)
	}
	log.SetFormatter(&log.TextFormatter{
		DisableColors:  false,
		DisableQuote:   true,
		FullTimestamp:  true,
		DisableSorting: false,
	})
	log.Info("Starting dummy server...")
	app.StartServer(port)
}
