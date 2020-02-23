package main

import (
    "fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/gotostn", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("request received")
        w.Write([]byte("Hello world"))
    })
    http.ListenAndServe("localhost:3000", mux)
}