package server

import (
	"database/sql"
	"github.com/Kirillznkv/new_year/api/config"
	"github.com/Kirillznkv/new_year/api/internal/store"
	_ "github.com/lib/pq"
	"net/http"
)

func Start() error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := store.New(db)
	srv := NewServer(store)

	return http.ListenAndServe(":8080", srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
