package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	Db *sql.DB
}

func New() *Store {
	return &Store{}
}

func (st *Store) Open() error {
	db, err := sql.Open("postgres", "host=localhost dbname=go_sql sslmode=disable")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	st.Db = db

	return nil
}

func (st *Store) Close() {
	if err := st.Db.Close(); err != nil {
		panic(err.Error())
	}
}
