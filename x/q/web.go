// q bert infinte
// okaq nano game
// AQ <aq@okaq.com>
// 2021-02-18
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "infinity.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("infinity live localhost:8080")
}

func InfinityHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", InfinityHandler)
	http.ListenAndServe(":8080", nil)
}


