package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello health!")
	})

	log.Printf("Starting a server at port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
