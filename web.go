// okaq.github.io
// dev web server
package main

import (
    "fmt"
    "net/http"
)

const (
    HTML = "index.html"
    HANA2D = "hana2d.js"
    OXY = "oxy.js"
    MJ2 = "mj2.js"
    PORT = ":8008"
    PINO = "pino.html"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, HTML)
    })
    http.HandleFunc("/hana2d.js", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, HANA2D)
    })
    http.HandleFunc("/oxy.js", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, OXY)
    })
    http.HandleFunc("/mj2.js", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, MJ2)
    })
    http.HandleFunc("/pino", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        http.ServeFile(w, r, PINO)
    })
 
    fmt.Println("Starting okaq.github.io dev web server on localhost" + PORT)
    http.ListenAndServe(PORT, nil)
}
