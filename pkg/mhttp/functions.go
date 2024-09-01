package mhttp

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}
var staticDir string

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, staticDir)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func hiString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}
