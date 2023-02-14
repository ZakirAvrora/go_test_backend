package main

import (
	"ZakirAvrora/go_test_backend/service1/config"
	"ZakirAvrora/go_test_backend/service1/db"
	"ZakirAvrora/go_test_backend/service1/internal/app"
	"ZakirAvrora/go_test_backend/service1/internal/store"
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	conf := config.NewConfig(".env")

	client, err := db.Init(conf.Database)

	if err != nil {
		log.Fatalln(err)
	}

	collection := client.Database("testing").Collection("users")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	ctx := context.TODO()
	s := store.New(collection, ctx)

	App := app.New(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/create-user", App.PostUser)
	r.Get("/get-user/{email}", App.GetUser)

	log.Println("Server started at :8081")
	log.Fatalln(http.ListenAndServe(":8081", r))
}
