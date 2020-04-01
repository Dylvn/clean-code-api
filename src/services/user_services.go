package services

import (
	"strings"

	"github.com/Dylvn/dragonfire-api/src/models/users"
)

var UsersService usersServiceInterface = &usersService{}

type usersService struct{}

type usersServiceInterface interface {
	StoreUser(u users.User) (*users.User, error)
	GetUser(userID int64) (*users.User, error)
	UpdateUser(u users.User) (*users.User, error)
	DeleteUser(u users.User) error
}

func (us *usersService) StoreUser(u users.User) (*users.User, error) {
	var err error

	if err = u.Validate(); err != nil {
		return nil, err
	}

	// TODO brypt the password
	if err = u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (us *usersService) GetUser(userID int64) (*users.User, error) {
	u := &users.User{Id: userID}
	if err := u.Get(); err != nil {
		return nil, err
	}

	return u, nil
}

func (us *usersService) UpdateUser(u users.User) (*users.User, error) {
	var err error
	var current = &users.User{Id: u.Id}

	if err = current.Get(); err != nil {
		return nil, err
	}

	current = handleUpdateUser(*current, u)

	if err = current.Validate(); err != nil {
		return nil, err
	}

	if err = current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (us *usersService) DeleteUser(u users.User) error {
	if err := u.Delete(); err != nil {
		return err
	}

	return nil
}

func handleUpdateUser(current, new users.User) *users.User {
	if new.Email = strings.TrimSpace(new.Email); new.Email != "" {
		current.Email = new.Email
	}
	if new.Password = strings.TrimSpace(new.Password); new.Password != "" {
		// TODO bcrypt the password
		current.Password = new.Password
	}
	if new.Fname = strings.TrimSpace(new.Fname); new.Fname != "" {
		current.Fname = new.Fname
	}
	if new.Lname = strings.TrimSpace(new.Lname); new.Lname != "" {
		current.Lname = new.Lname
	}

	return &current
}
