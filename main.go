package main

import (
	"letsgo/routes"
	"log"
	"net/http"
)

//PORT ...
const PORT = ":4000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.HomePage)
	mux.HandleFunc("/snippet", routes.ShowSnippet)
	mux.HandleFunc("/snippet/create", routes.SnippetCreate)

	log.Printf("Starting server on %s", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}
