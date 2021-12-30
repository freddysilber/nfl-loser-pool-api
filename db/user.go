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
		// err := scanUser(user, rows)
		err := rows.Scan(
			&user.Id,
			&user.Username,
			// &user.Email,
			// &user.FirstName,
			// &user.LastName,
			&user.Password,
			// &user.TokenHash,
			// &user.CreatedAt,
			// &user.UpdatedAt,
		)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) GetUserById(userId int) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE id = $1;`
	row := db.Conn.QueryRow(query, userId)
	switch err := row.Scan(
		&user.Id,
		&user.Username,
		// &user.FirstName,
		// &user.LastName,
		&user.Password,
		// &user.TokenHash,
		// &user.CreatedAt,
		// &user.UpdatedAt,
	); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
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
	// user.Password = GeneratehashPassword(user.Password)
	hashedPassword, err := hashAndSaltPassword([]byte(user.Password))
	log.Println("hashedPassword", hashedPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	// user.TokenHash = GenerateJWT(user.Username)

	var id int
	// var name string
	// var username string
	// var password string
	// var roles string[]
	// var createdAt string

	query := `INSERT INTO users (
		name,
		username, 
		password
	) VALUES ($1, $2, $3) RETURNING id`

	log.Println("Query", query)
	// query := `INSERT INTO users (
	// 	username, 
	// 	password, 
	// 	token_hash, 
	// 	first_name, 
	// 	last_name
	// ) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	err = db.Conn.QueryRow(
		query,
		user.Name,
		user.Username,
		user.Password,
		// user.TokenHash,
		// user.FirstName,
		// user.LastName,
	).Scan(&id)
	// ).Scan(&id, &name, &username, &password)

	log.Println(err);
	if err != nil {
		return err
	}
	user.Id = id
	// user.CreatedAt = createdAt
	return nil
}

func (db Database) GetUserByUsernameAndPassword(user *models.User) (*models.User, error) {
	log.Println("User username --> ", user.Username)
	query := `SELECT id, username, password FROM users WHERE username = $1`
	row := db.Conn.QueryRow(query, user.Username)
	log.Println("Queried User", row)
	switch err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
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

// generate a hashed-and-salted password from plain-text password. return value can be stored in db
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func hashAndSaltPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
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

// func scanUser(user models.User, rows *sql.Rows) error {
// 	return rows.Scan(
// 		&user.ID,
// 		&user.Username,
// 		// &user.Email,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.Password,
// 		&user.TokenHash,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	);
// }