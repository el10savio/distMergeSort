package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/distMergeSort/handlers"
)

const (
	// PORT defines the distMergeSort
	// node server port
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started distMergeSort node server")

	http.ListenAndServe(":"+PORT, r)
}
