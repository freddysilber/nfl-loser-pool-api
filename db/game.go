package db

import (
	"database/sql"

	"github.com/freddysilber/nfl-loser-pool-api/models"
)

func (db Database) AddGame(game *models.Game) error {
	var id int
	var ownerId int
	var name string
	var description string
	var createdAt string
	err := db.Conn.QueryRow(
		`INSERT INTO games (ownerId, name, description) VALUES ($1, $2, $3) RETURNING id, name, ownerId, description, created_at`,
		game.OwnerId,
		game.Name,
		game.Description,
	).Scan(
		&id,
		&name,
		&ownerId,
		&description,
		&createdAt,
	)
	if err != nil {
		return err
	}
	game.Id = id
	game.CreatedAt = createdAt
	return nil
}

func (db Database) GetAllGames() (*models.GameList, error) {
	list := &models.GameList{}
	rows, err := db.Conn.Query("SELECT * FROM games ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var game models.Game
		err := rows.Scan(&game.Id, &game.Name, &game.Description, &game.OwnerId, &game.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Games = append(list.Games, game)
	}
	return list, nil
}

func (db Database) DeleteGame(gameId int) error {
	_, err := db.Conn.Exec(`DELETE FROM games WHERE id = $1;`, gameId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
