package models

import (
	"database/sql"
	"time"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type Event struct {
	ID          int64 `json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	Location    string `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID		int64 `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := db.ExecuteQuery(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return utils.HandleError(err)
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return mapMultipleEvents(rows)
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	return mapSingleEvent(row)
}

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	_, err := db.ExecuteQuery(query, e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	_, err := db.ExecuteQuery(query, e.ID)
	return err
}

func (e *Event) Register(userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	_, err := db.ExecuteQuery(query, e.ID, userID)
	return err
}

func (e *Event) CancelRegistration(userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	_, err := db.ExecuteQuery(query, e.ID, userID)
	return err
}

func mapSingleEvent(row *sql.Row) (*Event, error) {
	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
		return nil, utils.HandleError(err)
	}
	return &event, nil
}

func mapMultipleEvents(rows *sql.Rows) ([]Event, error) {
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, utils.HandleError(err)
		}
		events = append(events, event)
	}
	return events, nil
}
