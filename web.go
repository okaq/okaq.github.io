// okaq.github.io
// dev web server
package main

import (
    "fmt"
    "net/http"
)

const (
    HTML = "index.html"
    PORT = ":8008"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, HTML)
    })
    fmt.Println("Starting okaq.github.io dev web server on localhost" + PORT)
    http.ListenAndServe(PORT, nil)
}
