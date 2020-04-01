package users

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"testing"
	"time"
)

var user = User{
	Id:       1,
	Email:    uniqid() + "@test.com",
	Fname:    "Foo",
	Lname:    "Bar",
	Password: "password",
}

func TestSaveUser(t *testing.T) {
	if err := user.Save(); err != nil {
		t.Errorf("The user should be created in db. err=%v \n", err)
	}
}

func TestGetUser(t *testing.T) {
	if err := user.Get(); err != nil {
		t.Errorf("The user should be getting from db. err=%v \n", err)
	}
}

func TestUpdateUser(t *testing.T) {
	if err := user.Update(); err != nil {
		t.Errorf("The user should be updated. err=%v \n", err)
	}
}

func uniqid() string {
	s10 := strconv.FormatInt(time.Now().Unix(), 10)
	hash := md5.New()
	hash.Write([]byte(s10))
	return hex.EncodeToString(hash.Sum(nil))
}
