package main

import (
	"ZakirAvrora/go_test_backend/service2/internals/app"
	"ZakirAvrora/go_test_backend/service2/internals/generator"
	"log"
	"net/http"
	"time"
)

func main() {

	gen := generator.New()

	app := app.New(gen)

	mux := http.NewServeMux()
	mux.HandleFunc("/generate-salt", app.Handle)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		WriteTimeout: 1 * time.Second,
	}

	log.Println("Server started at :8000")
	log.Fatalln(srv.ListenAndServe())
}
