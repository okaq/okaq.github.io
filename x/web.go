// xonix full screen
// okaq nano game
// AQ <aq@okaq.com>
// 2021-02-01
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "xonix.html"
	BOPS = "bops.html"
	FULL = "full.html"
	PLAY = "play.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("xonix live localhost:8080")
}

func XonixHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,INDEX)
}

func BopsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,BOPS)
}

func FullHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,FULL)
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,PLAY)
}

func main() {
	motd()
	http.HandleFunc("/", XonixHandler)
	http.HandleFunc("/b", BopsHandler)
	http.HandleFunc("/c", FullHandler)
	http.HandleFunc("/d", PlayHandler)
	http.ListenAndServe(":8080", nil)
}


