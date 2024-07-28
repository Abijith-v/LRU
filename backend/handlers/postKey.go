package handlers

import (
    "encoding/json"
    "net/http"
    "backend/lru"
)

func PostHandler(cache *lru.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var parameters map[string]string
        if err := json.NewDecoder(r.Body).Decode(&parameters); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        key, keyPresent := parameters["key"] // checking if the key is present
        value, valuePresent := parameters["value"] // checking if the value is present
        if !keyPresent || !valuePresent {
            http.Error(w, "Key or value missing", http.StatusBadRequest)
            return
        }
        cache.Put(key, value)
		response := map[string]string{
            "message": "Key has been added",
            "key":     key,
            "value":   value,
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(response)
    }
}
