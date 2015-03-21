package redisclient

import (
    "github.com/fzzy/radix/redis"
    "fmt"
)

func Connection() *redis.Client {
    client, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("Redis connection error. ", err)
    }
    return client
}
