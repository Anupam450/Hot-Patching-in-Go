package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting V2 server on :8080 (bug fixed!)")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from v2 (bug fixed!)")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
