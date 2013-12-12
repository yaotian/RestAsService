package main

import (
    "github.com/hoisie/web"
)


func service(val string) string { return "hello " + val } 

func main() {
    web.Get("/(.*)", service)
    web.Run("0.0.0.0:9999")
}
