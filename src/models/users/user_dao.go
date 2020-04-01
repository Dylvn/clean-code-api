package users

import (
	"github.com/Dylvn/dragonfire-api/src/config/db"
)

var (
	createUserQuery = "INSERT INTO users (email, password, fname, lname) VALUES ($1, $2, $3, $4) RETURNING id"
	getUserQuery    = "SELECT id, email, password, fname, lname FROM users WHERE id=$1"
	updateUserQuery = "UPDATE users SET email=$1, fname=$2, lname=$3, password=$4 WHERE id=$5"
	deleteUserQuery = "DELETE FROM users WHERE id=$1"
)

// Save a user in DB.
func (u *User) Save() error {
	stmt, err := db.DB.Prepare(createUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(u.Email, u.Password, u.Fname, u.Lname).Scan(&id)
	if err != nil {
		return err
	}

	u.Id = id

	return nil
}

// Get a specific user by his ID.
func (u *User) Get() error {
	stmt, err := db.DB.Prepare(getUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r := stmt.QueryRow(u.Id)
	if err = r.Scan(&u.Id, &u.Email, &u.Password, &u.Fname, &u.Lname); err != nil {
		return err
	}

	return nil
}

// Update the user in DB.
func (u *User) Update() error {
	stmt, err := db.DB.Prepare(updateUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Email, u.Password, u.Fname, u.Lname, u.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Delete() error {
	stmt, err := db.DB.Prepare(deleteUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id)
	if err != nil {
		return err
	}

	return nil
}
