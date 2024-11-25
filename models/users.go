package models

import (
	"errors"
	"example.com/event-booking-backend-app/db"
	"example.com/event-booking-backend-app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil
	}
	r, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userID, err := r.LastInsertId()
	u.ID = userID
	return err
}

func (u *User) ValidateCredentials() error {
	query := "select id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
