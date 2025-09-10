package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Build a hyperlink for dev
	url := fmt.Sprintf("http://localhost%s", *addr)

	log.Printf("Starting server on %s", *addr)
	// Log the URL
	log.Printf("Local:   \033[1;36m%s\033[0m\n", url)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
