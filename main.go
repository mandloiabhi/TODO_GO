package main

import (
	"log"
	"net/http"
	

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main(){

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	
	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

	log.Printf("Serving on port: 8080" )
	v1Router.Get("/healthz", handlerReadiness)
	router.Mount("/v1", v1Router)
	log.Fatal(srv.ListenAndServe())



}