package db

import (
	"log"

	"github.com/freddysilber/nfl-looser-pool-api/models"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM users ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next(){
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) AddUser(user *models.User) error {
	log.Println(user)
	var id int
	var createdAt string
	query := `INSERT INTO users (username) VALUES ($1) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, user.Username).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	user.ID = id
	user.CreatedAt = createdAt
	return nil
}
