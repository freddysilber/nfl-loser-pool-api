package db

import (
	"database/sql"

	"github.com/freddysilber/nfl-loser-pool-api/models"
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
			&user.Id,
			&user.Username,
			&user.Password,
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
	row := db.Conn.QueryRow(
		`SELECT * FROM users WHERE id = $1;`,
		userId,
	)
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

func (db Database) DeleteUser(userId int) error {
	_, err := db.Conn.Exec(`DELETE FROM users WHERE id = $1;`, userId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) NewUser(user *models.User) error {
	hashedPassword, err := hashAndSaltPassword([]byte(user.Password))
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	var id int

	query := `INSERT INTO users (
		name,
		username, 
		password
	) VALUES ($1, $2, $3) RETURNING id`

	err = db.Conn.QueryRow(
		query,
		user.Name,
		user.Username,
		user.Password,
	).Scan(&id)

	if err != nil {
		return err
	}
	user.Id = id
	return nil
}

func (db Database) GetUserByIdUsernameAndPassword(user *models.User) (*models.User, error) {
	row := db.Conn.QueryRow(
		`SELECT id, username, name, password FROM users WHERE id = $1 AND username = $2 AND password = $3`,
		user.Id,
		user.Username,
		user.Password,
	)
	switch err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Name,
		&user.Password,
	); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}

func (db Database) GetUserByUsername(user *models.User) (*models.User, error) {
	row := db.Conn.QueryRow(
		`SELECT id, username, name, password FROM users WHERE username = $1`,
		user.Username,
	)
	switch err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Name,
		&user.Password,
	); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}

func (db Database) GetGamesByUser(userId int) (*models.GameList, error) {
	list := &models.GameList{}
	games, err := db.Conn.Query("SELECT * FROM games WHERE ownerId = $1 ORDER BY name DESC", userId)
	if err != nil {
		return list, err
	}
	for games.Next() {
		var game models.Game
		err := games.Scan(&game.Id, &game.Name, &game.Description, &game.OwnerId, &game.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Games = append(list.Games, game)
	}
	return list, nil
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
