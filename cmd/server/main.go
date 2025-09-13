package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"plugin"
	"sync/atomic"
)

var currentHandler atomic.Value

// root response structure
type Endpoint struct {
	Path string `json:"path"`
	Desc string `json:"desc"`
}

func main() {
	log.Println("Starting server on :8080")

	// initial V1 handler with only /add
	currentHandler.Store(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			fmt.Println("List functions Endpoint hit")
			endpoints := []Endpoint{
				{Path: "/add", Desc: "Add two numbers"},
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

		default:
			http.NotFound(w, r)
		}
	}))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := currentHandler.Load().(http.Handler)
		handler.ServeHTTP(w, r)
	})

	// // simulate bug detection after 10s
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	log.Println("BUG detected: missing features, need to patch!")
	// }()

	// admin endpoint to hot patch
	http.HandleFunc("/admin/patch", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Loading v2.so for hot patch...")
		p, err := plugin.Open("v2.so")
		if err != nil {
			http.Error(w, "Failed to open plugin: "+err.Error(), 500)
			return
		}
		sym, err := p.Lookup("Handler")
		if err != nil {
			http.Error(w, "Failed to find Handler symbol: "+err.Error(), 500)
			return
		}
		handler, ok := sym.(func(http.ResponseWriter, *http.Request))
		if !ok {
			http.Error(w, "Invalid handler signature", 500)
			return
		}
		currentHandler.Store(http.HandlerFunc(handler))
		log.Println("Hot patch applied! Now serving V2")
		fmt.Fprintln(w, "Patched to V2")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
