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
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("xonix live localhost:8080")
}

func XonixHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", XonixHandler)
	http.ListenAndServe(":8080", nil)
}


