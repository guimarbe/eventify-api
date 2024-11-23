package models

import (
	"database/sql"
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64 `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return utils.HandleError(err)
	}

	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	result, err := db.ExecuteQuery(query, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return utils.HandleError(err)
	}

	u.ID = userID
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := mapUser(row, &u.ID, &retrievedPassword)
	if err != nil {
		return utils.HandleError(err)
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials are invalid")
	}

	return nil
}

func mapUser(row *sql.Row, dest ...interface{}) error {
	err := row.Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}
