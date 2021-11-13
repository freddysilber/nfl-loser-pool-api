package db

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/freddysilber/nfl-looser-pool-api/models"
	"golang.org/x/crypto/bcrypt"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM users ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, 
			&user.Username, 
			// &user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.TokenHash,
		)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) AddUser(user *models.User) error {
	hashedAndSalted, err := GeneratehashPassword(user.Password)
	user.Password = hashedAndSalted

	if err != nil {
		log.Fatalln("error in password hash")
	}
	
	user.TokenHash, err = GenerateJWT(user.Username)

	if err != nil {
		log.Fatalln("error creating JWT")
	}

	var id int
	var createdAt string
	query := `INSERT INTO users (
		username, 
		password, 
		token_hash, 
		first_name, 
		last_name
	) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	err = db.Conn.QueryRow(
		query, 
		user.Username, 
		user.Password, 
		user.TokenHash, 
		user.FirstName, 
		user.LastName,
		).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	user.ID = id
	user.CreatedAt = createdAt
	return nil
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

var secretkey = "mySuperSecretKey" // put this in the .env file
func GenerateJWT(username string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
