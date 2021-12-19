package db

import (
	"database/sql"
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

func (db Database) DeleteUser(userId int) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.Conn.Exec(query, userId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) SignUp(user *models.User) error {
	user.Password = GeneratehashPassword(user.Password)
	user.TokenHash = GenerateJWT(user.Username)

	var id int
	var createdAt string

	query := `INSERT INTO users (
		username, 
		password, 
		token_hash, 
		first_name, 
		last_name
	) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	err := db.Conn.QueryRow(
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

func (db Database) GetUserByUsernameAndPassword(user *models.User) (*models.User, error) {
	log.Println("User username --> ", user.Username)
	query := `SELECT * FROM users WHERE username = $1`
	row := db.Conn.QueryRow(query, user.Username)
	log.Println("Queried User", row)
	switch err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.TokenHash,
	); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}

// HashPassword is used to encrypt the password before it is stored in the DB
func GeneratehashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

var secretkey = "mySuperSecretKey" // put this in the .env file
func GenerateJWT(username string) string {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return ""
	}
	return tokenString
}

// VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "login or passowrd is incorrect"
		// msg = fmt.Sprintf("login or passowrd is incorrect")
		// msg = fmt.Sprintf("login or passowrd is incorrect %d", 100)
		check = false
	}

	return check, msg
}
