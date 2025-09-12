package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Endpoint struct {
	Path string `json:"path"`
	Desc string `json:"desc"`
}

// Handler is the new patched version
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		endpoints := []Endpoint{
			{Path: "/add", Desc: "Add two numbers"},
			{Path: "/subtract", Desc: "Subtract two numbers"},
			{Path: "/multiply", Desc: "Multiply two numbers"},
			{Path: "/divide", Desc: "Divide two numbers"},
		}
		json.NewEncoder(w).Encode(endpoints)

	case "/add":
		fmt.Println("Add handler hit")
		var a, b int
		_, err1 := fmt.Sscanf(r.URL.Query().Get("a"), "%d", &a)
		_, err2 := fmt.Sscanf(r.URL.Query().Get("b"), "%d", &b)
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid parameters", 400)
			return
		}
		fmt.Fprintf(w, "%d\n", a+b)

	case "/subtract":
		fmt.Println("Subtract handler hit")
		var a, b int
		_, err1 := fmt.Sscanf(r.URL.Query().Get("a"), "%d", &a)
		_, err2 := fmt.Sscanf(r.URL.Query().Get("b"), "%d", &b)
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid parameters", 400)
			return
		}
		fmt.Fprintf(w, "%d\n", a-b)

	case "/multiply":
		fmt.Println("Multiply handler hit")
		var a, b int
		_, err1 := fmt.Sscanf(r.URL.Query().Get("a"), "%d", &a)
		_, err2 := fmt.Sscanf(r.URL.Query().Get("b"), "%d", &b)
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid parameters", 400)
			return
		}
		fmt.Fprintf(w, "%d\n", a*b)

	case "/divide":
		fmt.Println("Divide handler hit")
		var a, b int
		_, err1 := fmt.Sscanf(r.URL.Query().Get("a"), "%d", &a)
		_, err2 := fmt.Sscanf(r.URL.Query().Get("b"), "%d", &b)
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid parameters", 400)
			return
		}
		if b == 0 {
			http.Error(w, "Division by zero", 400)
			return
		}
		fmt.Fprintf(w, "%d\n", a/b)

	default:
		http.NotFound(w, r)
	}
}
