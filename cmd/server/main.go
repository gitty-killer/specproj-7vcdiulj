package main

import (
    "log"
    "net/http"

    "'$name'/internal/handler"
    "'$name'/internal/store"
)

func main() {
    api := &handler.API{Store: store.New()}
    log.Println("listening on :3000")
    http.ListenAndServe(":3000", api)
}
