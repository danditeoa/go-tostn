package main

import (
    "net/http"

    "go-tostn/controller"
)

func main() {
    mux := controller.Register()
    http.ListenAndServe(":3000", mux)
}