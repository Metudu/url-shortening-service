package server

import (
	"net/http"

	"github.com/Metudu/url-shortening-service/api"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func routes() *mux.Router {
	router.HandleFunc("/shorten", api.CreateShortenedURL).Methods(http.MethodPost)
	router.HandleFunc("/shorten/{shortCode}", api.RetrieveOriginalURL).Methods(http.MethodGet)
	router.HandleFunc("/shorten/{shortCode}", api.UpdateURL).Methods(http.MethodPut)
	router.HandleFunc("/shorten/{shortCode}", api.DeleteURL).Methods(http.MethodDelete)
	router.HandleFunc("/shorten/{shortCode}/stats", api.GetStats).Methods(http.MethodGet)
	return router
}