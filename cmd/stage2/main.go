package main

import (
	"fmt"
	"log"
	"net/http"
	"plugin"
	"sync/atomic"
	"time"
)

var currentHandler atomic.Value

func main() {
	log.Println("🚀 Starting Stage2 server on :8080 with V1")

	// initial V1 handler
	currentHandler.Store(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from v1 (BUGGY!)")
	}))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := currentHandler.Load().(http.Handler)
		handler.ServeHTTP(w, r)
	})

	// simulate bug detection after 10s
	go func() {
		time.Sleep(10 * time.Second)
		log.Println("🐞 BUG detected: output incorrect, need to patch!")
	}()

	// admin endpoint to hot patch
	http.HandleFunc("/admin/patch", func(w http.ResponseWriter, r *http.Request) {
		log.Println("🔄 Loading v2.so for hot patch...")
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
		log.Println("✅ Hot patch applied! Now serving V2")
		fmt.Fprintln(w, "Patched to V2")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
