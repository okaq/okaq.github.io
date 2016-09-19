// okaq.github.io/anemonie.html
// dev web server
package main

import (
    "fmt"
    "net/http"
)

const (
    AN = "anemone.html"
    PORT = ":8080"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, AN)
    })
    fmt.Println("starting okaq.github.io/anemone.html dev web server on localhost" + PORT)
    http.ListenAndServe(PORT, nil)
}
