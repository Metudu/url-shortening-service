package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Metudu/url-shortening-service/db"
	"github.com/gorilla/mux"
)

type data struct {
	ID          *int       `json:"id,omitempty"`
	URL         *string    `json:"url,omitempty"`
	AccessCount *int       `json:"accesscount,omitempty"`
	ShortCode   *string    `json:"shortcode,omitempty"`
	CreatedAt   *time.Time `json:"createdat,omitempty"`
	UpdatedAt   *time.Time `json:"updatedat,omitempty"`
}

func printError(w http.ResponseWriter, method int, message string) {
	log.Println("[ERR] - " + message)
	w.WriteHeader(method)
}

func CreateShortenedURL(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB()
	if err := database.Ping(); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when pinging the database: %v", err.Error()))
		return
	}

	var payload data
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		printError(w, http.StatusBadRequest, fmt.Sprintf("An error occured when decoding the json: %v", err.Error()))
		return
	}

	if payload.URL == nil {
		printError(w, http.StatusBadRequest, "An error occured, the url value is null")
		return
	}

	if err := db.Append(database, *payload.URL); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when inserting the values to the database: %v", err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func RetrieveOriginalURL(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB()
	if err := database.Ping(); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when pinging the database: %v", err.Error()))
		return
	}

	query, err := db.GetByShortCode(database, mux.Vars(r)["shortCode"])
	if err != nil {
		printError(w, http.StatusNotFound, fmt.Sprintf("An error occured when getting the url by shortcode: %v", err.Error()))
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err := enc.Encode(*query); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when encoding the json: %v", err.Error()))
		return
	}
}

func UpdateURL(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB()
	if err := database.Ping(); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when pinging the database: %v", err.Error()))
		return
	}

	query, err := db.GetByShortCode(database, mux.Vars(r)["shortCode"])
	if err != nil {
		printError(w, http.StatusNotFound, fmt.Sprintf("An error occured when getting the url by shortcode: %v", err.Error()))
		return
	}
	var data data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when decoding the json data: %v", err.Error()))
		return
	}

	query.URL = data.URL

	if err := db.Update(database, query); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when updating the url: %v", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB()
	if err := database.Ping(); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when pinging the database: %v", err.Error()))
		return
	}

	if err := db.Delete(database, mux.Vars(r)["shortCode"]); err != nil {
		printError(w, http.StatusNotFound, fmt.Sprintf("An error occured when deleting the url: %v", err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetStats(w http.ResponseWriter, r *http.Request) {
    database := db.GetDB()
	if err := database.Ping(); err != nil {
		printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when pinging the database: %v", err.Error()))
		return
	}
	data, err := db.GetStats(database, mux.Vars(r)["shortCode"])
	if err != nil {
		printError(w, http.StatusNotFound, fmt.Sprintf("An error occured when getting the url by shortcode: %v", err.Error()))
		return
	}

    enc := json.NewEncoder(w)
    enc.SetIndent("", "    ")
    if err := enc.Encode(data); err != nil {
        printError(w, http.StatusInternalServerError, fmt.Sprintf("An error occured when encoding the json: %v", err.Error()))
        return
    }
}
