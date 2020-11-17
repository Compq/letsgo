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

//CreateSnippet ...
func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create Sinppet"))
}
