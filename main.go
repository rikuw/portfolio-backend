package main

import (
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	// If trying to get home page with other than GET request, serve and error
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	http.ServeFile(w, r, "public/index.html")
}

func main() {
	http.HandleFunc("/", serveHome)

	// Serve Javascript and CSS files
	fs := http.FileServer(http.Dir("public/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":80", nil))
}
