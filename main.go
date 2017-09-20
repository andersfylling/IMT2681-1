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

// Contains all the routes for the project
func main() {
	router := httprouter.New()

	// routes
	router.GET("/", Index)
	router.GET("/projectinfo/v1/github.com/:username/:repository", ParseGithubAPI)

	// start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
