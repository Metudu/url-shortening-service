package server

import (
	"log"
	"net/http"
	"time"
)


func InitializeServer() *http.Server {
	return &http.Server{
		Addr: ":8080",
		Handler: routes(),
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
	}
}

func StartServer(s *http.Server) {
	log.Println("Server is listening...")
	s.ListenAndServe()
}