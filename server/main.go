package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"html/template"
)

var mu sync.Mutex
var count int

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, World!")
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", echoString)
	http.HandleFunc("/count", counter)

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})



	log.Fatal(http.ListenAndServe(":8081", nil))

}
