package main

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Error("Port env var not set")
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Error("Failed to start http server", "err", err)
		os.Exit(1)
	}
}
