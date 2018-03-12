package main

import (
    "fmt"
    "net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprint(writer, "Hello World!")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":80", nil)
}
