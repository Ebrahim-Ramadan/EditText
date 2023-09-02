package main

import (
	"fmt"
	"net/http"
	"strings"
)

var editableContent = "<strong>Hello World!</strong>"

func main() {
	http.HandleFunc("/content", getContent) //endpoint that returns the current editable content as HTML.
	http.HandleFunc("/update", updateContent) // A POST endpoint that allows me to update the content. It expects the new content to be sent as a POST request parameter named content.
	http.ListenAndServe(":8080", nil)
}

func getContent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, editableContent)
}

func updateContent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	newContent := r.FormValue("content")
	if newContent == "" {
		http.Error(w, "Empty content not allowed", http.StatusBadRequest)
		return
	}

	editableContent = sanitizeHTML(newContent)
	fmt.Fprint(w, "Content updated successfully")
}

func sanitizeHTML(input string) string {
	// Basic HTML sanitization to prevent XSS attacks
	// I may want to use a more comprehensive library for this in a production environment
	// This is just a simple example
	// DO NOT use this in production without proper security measures
	return strings.ReplaceAll(input, "<script", "&lt;script")
}
