package models

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID        string    `json:"id" sql:"id"`
	Email     string    `json:"email" validate:"required" sql:"email"`
	Password  string    `json:"password" validate:"required" sql:"password"`
	Username  string    `json:"username" sql:"username"`
	FirstName  string    `json:"firstName" sql:"first_name"`
	LastName  string    `json:"lastName" sql:"last_name"`
	TokenHash string    `json:"tokenhash" sql:"tokenhash"`
	CreatedAt time.Time `json:"createdat" sql:"created_at"`
	UpdatedAt time.Time `json:"updatedat" sql:"updated_at"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (user *User) Bind(request *http.Request) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return fmt.Errorf("username is a required field")
	}
	return nil
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}