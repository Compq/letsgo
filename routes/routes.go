package routes

import "net/http"

//HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there"))
}

//ShowSnippet ...
func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet"))
}

//SnippetCreate ...
func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Snippet"))
}
