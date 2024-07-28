package main

import (
	"fmt"
    "net/http"
	"backend/routes"
	"backend/lru"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
)

type KeyValue struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

var store = make(map[string]string)

func main() {
	
	cache := lru.CreateLRUCache(5, 1*time.Minute) // capacity is 5 and ttl is 5 minutes

	r := chi.NewRouter()

	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Allow specific origin
        AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"}, // Allow specific methods
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        ExposedHeaders:   []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           300, // Max age for preflight requests
    })

    // Add middleware
    r.Use(middleware.Logger)
    r.Use(c.Handler)

	routes.NewRouter(r, cache)
	// starting the server
	fmt.Println("Running server at port : 8080")
    http.ListenAndServe(":8080", r)
}



