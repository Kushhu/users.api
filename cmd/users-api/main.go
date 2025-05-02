package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kushhu/users.api.git/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to users api"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("server initiated...")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server")
	}
}
