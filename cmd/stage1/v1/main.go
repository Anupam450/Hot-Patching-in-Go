package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting V1 server on :8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from v1")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
