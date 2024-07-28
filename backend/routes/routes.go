package routes

import (
    "github.com/go-chi/chi/v5"
    "backend/handlers"
    "backend/lru"
)

func NewRouter(r *chi.Mux, cache *lru.LRUCache) *chi.Mux {

    r.Get("/get", handlers.GetHandler(cache))
    r.Post("/put", handlers.PostHandler(cache))
    r.Delete("/delete", handlers.DeleteHandler(cache))

    return r
}
