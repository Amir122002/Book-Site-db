package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/create", Create).Methods(http.MethodPost)
	router.HandleFunc("/read/{page}", Read).Methods(http.MethodGet)
	router.HandleFunc("/read_id/{id}", ReadId).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", Update).Methods(http.MethodPut)
	router.HandleFunc("/delete/{id}", Delete).Methods(http.MethodDelete)

	return router
}
