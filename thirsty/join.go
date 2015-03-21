package thirsty

import (
    "net/http"
    "github.com/frodetbj/beeroclock-api/redisclient"
    "fmt"
    "encoding/json"
)

type ThirstyDude struct {
    Username string
}

func response(w http.ResponseWriter, statuscode int, payload interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(statuscode)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        return err
    }

    return nil
}

func Join(w http.ResponseWriter, r *http.Request) {
    var thirstydude ThirstyDude
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&thirstydude)
    if err != nil {
        panic(err)
    }

    if thirstydude.Username == "" {
        response(w, 400, "Username missing")
    } else {
        client := redisclient.Connection()
        err = client.Cmd("SADD", "thirstydudes", thirstydude).Err
        if err != nil {
            fmt.Println("Redis cmd fail ", err)
        }

        response(w, 200, thirstydude)
    }
}
