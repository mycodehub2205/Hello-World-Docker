package main

import (
	"log"
	"net/http"

	"github.com/synoa/helloworld"
)

func main() {
	gr := helloworld.NewGreeter("Hello")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/"):]
		w.Write([]byte(gr.Greet(name)))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
