package main

import (
    "net/http"
    "github.com/frodetbj/beeroclock-api/thirsty"
    "fmt"
)

func main() {
    http.HandleFunc("/thirsty",thirsty.Handler)
    fmt.Println(http.ListenAndServe(":8080", nil))
}
