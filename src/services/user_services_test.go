package services

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	u, err := UsersService.GetUser(1)
	if err != nil {
		t.Error(err)
	}

	if err = u.Validate(); err != nil {
		t.Error(err)
	}
}
