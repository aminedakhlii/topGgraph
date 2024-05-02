package api

import (
	"github.com/gorilla/mux"
)

// NewRouter creates and returns a new router.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/nodes/{id}", GetNodeHandler).Methods("GET")
	router.HandleFunc("/nodes", CreateNodeHandler).Methods("POST")

	return router
}
