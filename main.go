package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
    var addr string = ":8080"

    server := http.Server {
        Addr: addr,
        Handler: &myHandler{},
        ReadTimeout: 5 * time.Second,
    }

    fmt.Println("Server Listening on port", addr)

    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/tmp"] = Tmp
    mux["/"] = JsonAPI
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}

type Post struct {
    Title string `json:"Title"`
    Author string `json:"Author"`
    Text string `json:"Text"`
}

func JsonAPI(w http.ResponseWriter, r *http.Request) {
    // Example: add Post structs
    posts := []Post {}

    json.NewEncoder(w).Encode(posts)
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := mux[r.URL.String()];ok {
        h(w, r)
        return
    }
    io.WriteString(w, "URL: " + r.URL.String())
}

func Tmp(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "version 3")
}
