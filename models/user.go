package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID			int
	FirstName	string
	LastName	string
}

var (
	users	 []*User
	nextID	= 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("The user ID should be empty")
	}

	u.ID = nextID
	nextID++

	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("The user with ID '%v' not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("The user with ID '%v' not found", user.ID)
}

func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("The user with ID '%v' not found", id)
}
