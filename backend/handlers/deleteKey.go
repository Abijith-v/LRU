package handlers

import (
    "net/http"
    "backend/lru"
)

func DeleteHandler(cache *lru.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cache.DeleteLRU()
        w.WriteHeader(http.StatusNoContent)
    }
}
