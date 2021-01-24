package domain

import "errors"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

var userList = []User{
	{Id: 1, Name: "John Doe", Email: "john@example.org", Password: "secret"},
}

func GetUserByEmail(email string) (User, error) {
	for _, u := range userList {
		if u.Email == email {
			return u, nil
		}
	}

	return User{}, errors.New("user not found")
}
