package main

import (
	_"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")

	portstring := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string {"https://*", "http://*"},
		AllowedMethods: []string {"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string {"*"},
		ExposedHeaders: []string {"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router:= chi.NewRouter()

	v1Router.Get("/ready", handleReadiness)
	v1Router.Get("/err", handleErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portstring,
	}

	log.Printf("server starting on port: %v", portstring)
	err:=srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

} 
	


