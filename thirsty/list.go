package thirsty

import (
    "net/http"
    "github.com/frodetbj/beeroclock-api/redisclient"
    "github.com/frodetbj/beeroclock-api/utils"
)

func List(w http.ResponseWriter, r *http.Request) {
    client := redisclient.Connection()
    res := client.Cmd("SMEMBERS", "thirstydudes")

    if res.Err != nil {
        utils.Response(w, 500, "Redis fail")
        panic(res.Err)
    }

    var dudes []string
    for _, line := range res.Elems {
        str, _ := line.Str()
        dudes = append(dudes, str)
    }

    utils.Response(w, 200, dudes)
}
