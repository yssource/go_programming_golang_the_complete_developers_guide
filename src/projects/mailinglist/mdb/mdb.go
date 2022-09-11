package mdb

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type EmailEntry struct {
	Id          int64
	Email       string
	ConfirmedAt *time.Time
	OptOut      bool
}

func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE emails (
      id           INTEGER PRIMARY KEY,
      email        TEXT UNIQUE,
      confirmed_at INTEGER,
      opt_out      INTEGER
    );
  `)

	// NOTE: casting err type to sqlite3.Error type
	if sqlError, ok := err.(sqlite3.Error); ok {
		// NOTE: 1 means table already exists
		if sqlError.Code != 1 {
			log.Fatal(sqlError)
		}
	} else {
		log.Fatal(err)
	}
}

func EmailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	var (
		id          int64
		email       string
		confirmedAt int64
		optOut      bool
	)

	err := row.Scan(&id, &email, &confirmedAt, &optOut)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	t := time.Unix(confirmedAt, 0)

	return &EmailEntry{
		Id:          id,
		Email:       email,
		ConfirmedAt: &t,
		OptOut:      optOut,
	}, nil
}
