package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index smth
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func HandleGitHubRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	gh := NewRepo(ps.ByName("username"), ps.ByName("repository"))
	gh.SetRepoDetails(false)

	fmt.Fprintf(w, gh.getJSON())
}

// Contains all the routes for the project
func main() {
	router := httprouter.New()

	// routes
	router.GET("/", Index)
	router.GET("/projectinfo/v1/github.com/:username/:repository", HandleGitHubRequest)

	// start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
