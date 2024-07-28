package handlers

import (
    "encoding/json"
    "net/http"
    "backend/lru"
)

func GetHandler(cache *lru.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        
        keys := cache.Keys() // get all the keys
        if len(keys) == 0 {
            http.Error(w, "Cache is empty", http.StatusNotFound)
            return
        }

        // Get the least recently used key
        lruKey := keys[len(keys)-1]
        value, ok := cache.Get(lruKey)
        if !ok {
            http.Error(w, "LRU key has expired", http.StatusNotFound)
            return
        }

        response := map[string]string{"key": lruKey, "value": value}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}
