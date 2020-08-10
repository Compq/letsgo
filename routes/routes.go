package routes

import "net/http"

//HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hi from home!"))
}

//ShowSnippet ...
func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet"))
}

//SnippetCreate ...
func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Snippet"))
}
