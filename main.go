package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Error("Port env var not set")
		os.Exit(1)
	} else {
		log.Info("Starting server on port", "port", port)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	rootHandle := func(w http.ResponseWriter, r *http.Request) {
		rootPage := template.Must(template.ParseFiles("./site/index.html"))
		rootPage.Execute(w, nil)
	}

	http.HandleFunc("/", rootHandle)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Error("Failed to start http server", "err", err)
		os.Exit(1)
	}
}
