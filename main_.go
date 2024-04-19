package main

import (
	"net/http"
)

func main() {
	// Serve the HTML file for the root path
	http.HandleFunc("/", handleRoot)

	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Start the server
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML file
	http.ServeFile(w, r, "public/ExampleBlogPage.html")
}
