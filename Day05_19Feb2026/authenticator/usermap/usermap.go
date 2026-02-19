package usermap

import (
	"authenticator/user"
	"errors"
	"fmt"
)

type UserMap struct {
	UserMap map[string]*user.User
}

func New() *UserMap {
	return &UserMap{
		UserMap: make(map[string]*user.User),
	}
}

func (um *UserMap) Add(user *user.User) error {
	_, ok := um.UserMap[user.Username]
	if ok {
		return errors.New("User already exists!")
	}
	um.UserMap[user.Username] = user
	fmt.Println("User:", user.Username, "added successfully!")
	return nil
}

func (um *UserMap) Authenticate(username, password string) error {
	_, ok := um.UserMap[username]
	if !ok {
		return fmt.Errorf("User %s doesn't exist!", username)
	}

	if um.UserMap[username].Password != password {
		return errors.New("Wrong password!")
	}

	fmt.Println("User:", username, "logged in successfully!")
	return nil
}
