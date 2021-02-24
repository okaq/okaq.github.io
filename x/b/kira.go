// okaq dojo kira
// okaq bitmap sampler
// AQ <aq@okaq.com>
// 2021-02-18
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

const (
	INDEX = "kira.html"
	DATA = "kirad/"
)

var (
	// file paths
	P []string
	// json array of file paths
	J []byte
	// base64 bitmap data encoding
	B []string
	// final json file
	D []byte
	C uint64
)

func KiraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PathsHandler(w http.ResponseWriter, r *http.Request) {
	// image files paths json
	// embed data on INDEX first call
	// perhaps in html template json
	// or inline uri data
	fmt.Println(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(J)
}

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// serve by file path
	// array buffer data
	// use atomic counter to index
	// b0 := bytes.NewBuffer(r.Body)
	b0 := new(bytes.Buffer)
	b0.ReadFrom(r.Body)
	fmt.Println(b0.Bytes())
	// json marshal
	s0 := fmt.Sprintf("recieved %d bytes ok", b0.Len())
	b2, err := json.Marshal(b0.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b2))
	// increment
	i0 := atomic.LoadUint64(&C)
	fmt.Printf("Current file at index %d: %s\n", i0, P[i0])
	// set data
	B[i0] = string(b2)
	atomic.AddUint64(&C, 1)
	// when counter reaches len(P) 
	// stitch and write to disk
	n0 := uint64(len(P) - 1)
	if i0 >= n0 {
		fmt.Println(B)
		// done
		// json string
		// file path format: kirad/wand_1.png
		// get prefix for class name
		s2 := strings.Split(P[0], "/")
		s3 := strings.Split(s2[1], "_")
		s4 := fmt.Sprintf("%ss", s3[0])
		s5 := fmt.Sprintf("var %s = {\n", s4)
		s1 := new(bytes.Buffer)
		s1.WriteString(s5)
		for i, p0 := range P {
			s6 := strings.Split(p0, "/")
			s7 := strings.Split(s6[1], ".")
			s8 := fmt.Sprintf("\"%s\":%s,\n",s7[0],B[i])
			s1.WriteString(s8)
		}
		s10 := fmt.Sprintf("}\n\n")
		s1.WriteString(s10)
		s9 := fmt.Sprintf("%s.json", s4)
		ioutil.WriteFile(s9,s1.Bytes(),0666)
	}
	b1 := []byte(s0)
	w.Write(b1)
}

func files() {
	// generate list of file paths
	// sort in lexicographic order
	f0, err := ioutil.ReadDir(DATA)
	if err != nil {
		fmt.Println(err)
	}
	P = make([]string, len(f0))
	for i, f1 := range f0 {
		// fmt.Println(f1.Name())
		// P[i] = f1.Name()
		f2 := fmt.Sprintf("%s%s", DATA, f1.Name())
		P[i] = f2
	}
	fmt.Println(P)
	var err2 error
	J, err2 = json.Marshal(P)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(J))
	B = make([]string, len(f0))
}

func meta() {
	// load persistant log file
	// unmarshal to local object
	// launch goroutine to update
	// marshal json and write to new log
}

func motd() {
	t0 := time.Now()
	fmt.Println("okaq web serve on localhost:8080")
	r0 := rand.New(rand.NewSource(t0.UnixNano())).Uint32()
	fmt.Printf("id: %d.\ntime: %s.\n", r0, t0.String())
}

func main() {
	motd()
	meta()
	files()
	http.HandleFunc("/", KiraHandler)
	http.HandleFunc("/a", PathsHandler)
	http.HandleFunc("/b", FilesHandler)
	// static server
	fs := http.FileServer(http.Dir(DATA))
	http.Handle("/kirad/", http.StripPrefix("/kirad/", fs))
	http.ListenAndServe(":8080", nil)
}

// single data dir
// bat.png, bat.json, data.js

// persistent log
// write json metadata to filesystem

// test wands.json output
// 1868 bytes for 10 bitmaps (1024x1024)
// versus 300kB for 10 PNGs

// write json output to dedicated directory
// cache copies on github (no ignore)

}

