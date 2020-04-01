package users

import (
	"errors"
	"net/mail"
	"strings"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fname    string `json:"first_name"`
	Lname    string `json:"last_name"`
}

func (u *User) Validate() error {
	if err := u.validateEmail(); err != nil {
		return err
	}
	if u.Password = strings.TrimSpace(u.Password); u.Password == "" {
		return errors.New("The user password can't be empty")
	}
	if u.Fname = strings.TrimSpace(u.Fname); u.Fname == "" {
		return errors.New("The user first_name can't be empty")
	}
	if u.Lname = strings.TrimSpace(u.Lname); u.Lname == "" {
		return errors.New("The user last_name can't be empty")
	}

	return nil
}

func (u *User) validateEmail() error {
	if u.Email = strings.TrimSpace(u.Email); u.Email == "" {
		return errors.New("The user email can't be empty")
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}

	return nil
}
