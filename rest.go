package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	port := os.Getenv("PORT")
	vmux := mux.NewRouter()

	vmux.Handle("/api/", apiHandler{})
	vmux.HandleFunc("/", getindex)
	vmux.HandleFunc("/greets/{name}", greets)

	http.ListenAndServe(":"+port, vmux)
}

func getindex(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, "{\"content\":\"Welcome to the home page!\"}")
}

func greets(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	vars := mux.Vars(req)
	name := vars["name"]
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, "{\"content\":\"Hello %s \"}", name)

}
