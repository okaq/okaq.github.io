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

func main() {
	motd()
	http.HandleFunc("/", XonixHandler)
	http.HandleFunc("/b", BopsHandler)
	http.ListenAndServe(":8080", nil)
}


