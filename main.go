package main

import (
	"fmt"
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
	}))
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
	


