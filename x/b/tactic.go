// bacterial tactics
// okaq nano game
// AQ <aq@okaq.com>
// 2021-02-23
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "tactic.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("tactics live localhost:8080")
}

func TacticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,INDEX)
}

func NotoHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"noto_emoji_2.json")
}

func main() {
	motd()
	http.HandleFunc("/", TacticHandler)
	http.HandleFunc("/a", NotoHandler)
	http.ListenAndServe(":8080", nil)
}


