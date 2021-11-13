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
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) AddUser(user *models.User) error {
	var err error
	log.Println(user)
	log.Println("Password", user.Password)    // Hash and salt this
	log.Println("Token Hash", user.TokenHash) // Generate a new JWT
	user.Password, err = GeneratehashPassword(user.Password)
	log.Println("Hashed and Salted Password", user.Password) // Hash and salt this

	if err != nil {
		log.Fatalln("error in password hash")
	}
	user.TokenHash, err = GenerateJWT(user.Email)
	if err != nil {
		log.Fatalln("error creating JWT")
	}
	log.Println("Token Hash", user.TokenHash) // Generate a new JWT
	var id int
	var createdAt string
	query := `INSERT INTO users (username, email, password, token_hash) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err = db.Conn.QueryRow(query, user.Username, user.Email, user.Password, user.TokenHash).Scan(&id, &createdAt)
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
func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// func GenerateJWT(email, role string) (string, error) {
// 	var mySigningKey = []byte(secretkey)
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["email"] = email
// 	claims["role"] = role
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil
// }
