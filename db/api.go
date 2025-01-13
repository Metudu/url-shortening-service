package db

import (
	"database/sql"
	"errors"
	"math/rand/v2"
	"time"
)

type data struct {
	ID          *int       `json:"id,omitempty"`
	URL         *string    `json:"url,omitempty"`
	AccessCount *int       `json:"accessount,omitempty"`
	ShortCode   *string    `json:"shortcode,omitempty"`
	CreatedAt   *time.Time `json:"createdat,omitempty"`
	UpdatedAt   *time.Time `json:"updatedat,omitempty"`
}

func createShortCode() string {
	shortCode := ""
	characters := "abcdefghijklmnopqrstuvxyz1234567890"
	length := 6

	for i := 0; i < length; i++ {
		shortCode += string(characters[rand.IntN(len(characters))])
	}

	return shortCode
}

func Append(db *sql.DB, url string) error {
	_, err := db.Exec(`
        INSERT INTO "servicedb"(url, shortcode, accesscount, createdat, updatedat) VALUES ($1,$2,$3,$4,$5)`,
		url,
		createShortCode(),
		0,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func GetByShortCode(db *sql.DB, shortCode string) (*data, error) {
	var data data
	row := db.QueryRow(`SELECT id, url, shortcode, createdat, updatedat FROM servicedb WHERE shortcode = $1`, shortCode)
	if err := row.Scan(&data.ID, &data.URL, &data.ShortCode, &data.CreatedAt, &data.UpdatedAt); errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &data, nil
}

func Update(db *sql.DB, data *data) error {
	_, err := db.Exec(`UPDATE servicedb SET url = $1, updatedat = $2 WHERE shortcode = $3`, *data.URL, time.Now(), *data.ShortCode)
	if err != nil {
		return err
	}
	return nil
}

func Delete(db *sql.DB, shortCode string) error {
	_, err := db.Exec(`DELETE FROM servicedb WHERE shortcode = $1`, shortCode)
	if err != nil {
		return err
	}
	return nil
}

func updateAccessCount(db *sql.DB, shortCode string) error {
	_, err := db.Exec(`UPDATE servicedb SET accesscount = accesscount + 1 WHERE shortcode = $1`, shortCode)
	if err != nil {
		return err
	}
	return nil
}

func GetStats(db *sql.DB, shortCode string) (*data, error) {
	if err := updateAccessCount(db, shortCode); err != nil {
		return nil, err
	}

	var data data
	row := db.QueryRow(`SELECT id, url, accesscount, shortcode, createdat, updatedat FROM servicedb WHERE shortcode = $1`, shortCode)
	if err := row.Scan(&data.ID, &data.URL, &data.AccessCount, &data.ShortCode, &data.CreatedAt, &data.UpdatedAt); errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &data, nil
}