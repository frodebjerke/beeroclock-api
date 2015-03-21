package thirsty

import (
    "net/http"
    "github.com/frodetbj/beeroclock-api/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
        case "GET":
            List(w, r)
        case "POST":
            Join(w, r)
        default:
            utils.Response(w, 404, r.URL.Path)
    }
}
