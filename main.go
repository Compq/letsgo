package main

import (
	"letsgo/routes"
	"log"
	"net/http"
)

//PORT ...
const PORT = ":4000"

func main() {

	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/snippet", routes.ShowSnippet)
	http.HandleFunc("/snippet/create", routes.CreateSnippet)

	log.Printf("Starting server on %s", PORT)
	err := http.ListenAndServe(PORT, nil)
	log.Fatal(err)
}
