package handler

import (
    "encoding/json"
    "net/http"
    "strings"

    "'$name'/internal/store"
)

type API struct {
    Store *store.Store
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/health" {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
        return
    }

    if strings.HasPrefix(r.URL.Path, "/item/") && r.Method == http.MethodGet {
        key := strings.TrimPrefix(r.URL.Path, "/item/")
        if value, ok := a.Store.Get(key); ok {
            json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
            return
        }
        w.WriteHeader(http.StatusNotFound)
        return
    }

    if strings.HasPrefix(r.URL.Path, "/item/") && r.Method == http.MethodPost {
        key := strings.TrimPrefix(r.URL.Path, "/item/")
        buf := make([]byte, r.ContentLength)
        r.Body.Read(buf)
        a.Store.Set(key, string(buf))
        w.WriteHeader(http.StatusCreated)
        return
    }

    w.WriteHeader(http.StatusNotFound)
}
