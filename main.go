package main

import (
	"log"
	"net/http"

	"github.com/ab22/env"
)

type AppConfig struct {
	Env  string `env:"ENV" envDefault:"DEV"`
	Port int    `env:"PORT" envDefault:"1337"`
}

func main() {
	config := &AppConfig{}

	log.Println("Stating server...")

	env.Parse(config)
	log.Println("Server port:", config.Port)
	log.Println("App environment:", config.Env)

	templatesFilesHandler := http.FileServer(http.Dir("html/templates"))
	staticFilesHandler := noDirListing(http.FileServer(http.Dir("html/static/")))

	http.Handle("/", templatesFilesHandler)
	http.Handle("/static/", http.StripPrefix("/static/", staticFilesHandler))

	log.Println("Listening...")
	http.ListenAndServe(":1337", nil)
}
