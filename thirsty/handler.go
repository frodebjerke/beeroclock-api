package thirsty

import (
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

    if r.Method == "POST" {
        Join(w, r)
    }
}
