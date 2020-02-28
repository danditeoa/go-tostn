package main

import (
    "fmt"
    "log"
    "net/http"

    "go-tostn/controller"
)

func main() {
    mux := controller.Register()
    fmt.Println("Serving...")
    log.Fatal(http.ListenAndServe(":3000", mux))
}
