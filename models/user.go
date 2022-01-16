package models

import (
	"fmt"
	"net/http"
)

type User struct {
	// Id       int      `json:"id" sql:"id"`
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	// Roles     []string `json:"roles"`
	CreatedAt string   `json:"createdAt"`
}

// type User struct {
// 	ID        int    `json:"id" sql:"id"`
// 	// Email     string    `json:"email" validate:"required" sql:"email"`
// 	Password  string    `json:"password" validate:"required" sql:"password"`
// 	Username  string    `json:"username" sql:"username"`
// 	FirstName  string    `json:"firstName" sql:"first_name"`
// 	LastName  string    `json:"lastName" sql:"last_name"`
// 	TokenHash string    `json:"tokenhash" sql:"token_hash"`
// 	CreatedAt string `json:"createdat" sql:"created_at"`
// 	UpdatedAt string `json:"updatedat" sql:"updated_at"`
// 	// CreatedAt time.Time `json:"createdat" sql:"created_at"`
// 	// UpdatedAt time.Time `json:"updatedat" sql:"updated_at"`
// }

type UserList struct {
	Users []User `json:"users"`
}

func (user *User) Bind(request *http.Request) error {
	if user.Username == "" {
		// if user.Username == "" || user.Email == "" || user.Password == "" {
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
