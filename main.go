package main

import (
    "io"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", Tmp)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func Tmp(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "version 1")
}
