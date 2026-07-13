package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mdiktushar/REST-API-with-GoLang/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})
	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Fail to start server")
	}
}
