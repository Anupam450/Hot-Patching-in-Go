package main

import (
	"fmt"
	"net/http"
)

// Handler is the new patched version
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from v2 (fixed via hot patch!)")
}
