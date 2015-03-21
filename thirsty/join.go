package thirsty

import (
    "net/http"
    "github.com/frodetbj/beeroclock-api/redisclient"
    "github.com/frodetbj/beeroclock-api/utils"
    "fmt"
    "encoding/json"
)

func Join(w http.ResponseWriter, r *http.Request) {
    var thirstydude Dude
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&thirstydude)
    if err != nil {
        utils.Response(w, 400, err)
        panic(err)
    }

    if thirstydude.Username == "" {
        utils.Response(w, 400, "Username missing")
    } else {
        client := redisclient.Connection()
        err = client.Cmd("SADD", "thirstydudes", thirstydude.Username).Err
        if err != nil {
            fmt.Println("Redis cmd fail ", err)
        }

        utils.Response(w, 200, thirstydude)
    }
}
