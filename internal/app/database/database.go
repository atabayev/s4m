package database

import (
	"database/sql"
	"math/rand"
	"s4m/internal/app/model"
	"time"
)

type Database struct {
	connect *sql.DB
}

func New(driveName string, DSN string) (*Database, error) {
	conn, err := sql.Open(driveName, DSN)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return &Database{connect: conn}, nil
}

func (db *Database) InsertOne(event model.Event) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	_, err := db.connect.Exec("INSERT INTO events (eventID, eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?, ?)",
		10000+r.Int63n(89999), event.EventType, event.UserID, event.EvenTime, event.Payload)
	return err
}

func (db *Database) InsertMore(events []model.Event) error {
	tx, err := db.connect.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO events (eventID, eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, event := range events {
		_, err = stmt.Exec(10000+r.Int63n(89999), event.EventType, event.UserID, event.EvenTime, event.Payload)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (db *Database) SelectEvent(eventType string, startDate, finishDate time.Time) ([]model.Event, error) {
	rows, err := db.connect.Query("SELECT * FROM events WHERE eventType=? AND eventTime BETWEEN ? AND ?", eventType, startDate, finishDate)

	if err != nil {
		return nil, err
	}

	var e model.Event
	var events []model.Event
	for rows.Next() {
		if err = rows.Scan(
			&e.EventID,
			&e.EventType,
			&e.UserID,
			&e.EvenTime,
			&e.Payload,
		); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
