package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShahJabir/student-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api"))
	})
	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	fmt.Printf("Server Started %s \n", cfg.HTTPServer.Address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
