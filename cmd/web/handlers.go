package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet hanler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...\n", id)
	//w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Cache-Control", "public, max-age=3153600")
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")

		w.Header().Del("Cache-Control")

		w.Header().Get("Cache-Control")

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name":"Santi"}`))

		w.Header()["X-XSS-Protection"]= []string{"1; mode=block"}
		w.Header()["Date"] = nil

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
