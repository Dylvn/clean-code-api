package users

import "testing"

var u = &User{
	Email:    "test@test.com",
	Password: "test",
	Fname:    "test",
	Lname:    "test",
}

func TestValidateWithValidUser(t *testing.T) {
	if err := u.Validate(); err != nil {
		t.Errorf("A valid user should not return an error. err='%v'", err)
	}
}

func TestValidateUserWithEmptyEmail(t *testing.T) {
	u.Email = ""

	if err := u.Validate(); err == nil {
		t.Errorf("An error should be return if the user email is empty")
	}
}

func TestValidateUserWithEmptyPassword(t *testing.T) {
	u.Password = ""

	if err := u.Validate(); err == nil {
		t.Errorf("An error should be return if the user email is empty")
	}
}

func TestValidateUserWithEmptyFname(t *testing.T) {
	u.Fname = ""

	if err := u.Validate(); err == nil {
		t.Errorf("An error should be return if the user email is empty")
	}
}

func TestValidateUserWithEmptyLname(t *testing.T) {
	u.Lname = ""

	if err := u.Validate(); err == nil {
		t.Errorf("An error should be return if the user email is empty")
	}
}

func TestValidateUserWithInvalidEmail(t *testing.T) {
	u.Email = "testtest.com"

	if err := u.Validate(); err == nil {
		t.Errorf("An error should be return if the user email is empty")
	}
}
