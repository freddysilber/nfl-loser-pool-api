package db

import (
	"database/sql"
	"fmt"

	"github.com/freddysilber/nfl-loser-pool-api/models"
)

func (db Database) AddGame(game *models.Game, shareId string, gameId string) error {
	var ownerId string
	var name string
	var description string
	var createdAt string
	err := db.Conn.QueryRow(
		`
			INSERT INTO games (id, owner_id, name, description, share_id)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, name, owner_id, description, created_at, share_id
		`,
		gameId,
		game.OwnerId,
		game.Name,
		game.Description,
		shareId,
	).Scan(
		&gameId,
		&name,
		&ownerId,
		&description,
		&createdAt,
		&shareId,
	)
	if err != nil {
		return err
	}
	// TODO: fix this so we dont have to set these props to access them in the client
	game.Id = gameId
	game.CreatedAt = createdAt
	game.ShareId = shareId
	return nil
}

func (db Database) GetAllGames() (*models.GameList, error) {
	list := &models.GameList{}
	rows, err := db.Conn.Query(`
		SELECT *
		FROM games
		ORDER BY ID DESC
	`)

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var game models.Game

		// These 'Scan' calls work the best when they are in the exact order as the SQL tables
		err := rows.Scan(
			&game.Id,
			&game.Name,
			&game.Description,
			&game.ShareId,
			&game.OwnerId,
			&game.CreatedAt,
		)

		if err != nil {
			return list, err
		}
		
		list.Games = append(list.Games, game)
	}
	return list, nil
}

// TODO: only game owners should be able to delete their own games
func (db Database) DeleteGame(gameId string) error {
	// When we delete a game, we need to delete all the game players with it since they are in a one to many required relationship
	var err error
	_, err = db.Conn.Exec(
		`
			DELETE
			FROM players
			WHERE game_id = $1
		`,
		gameId,
	)
	if err != nil {
		return err
	}
	// Then delete the game record
	_, err = db.Conn.Exec(
		`
			DELETE
			FROM games
			WHERE id = $1;
		`,
		gameId,
	)

	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) GetGame(gameId string) (*models.Game, error) {
	game := &models.Game{}
	row := db.Conn.QueryRow(
		`
			SELECT * FROM games WHERE id = $1
		`,
		gameId,
	)
	switch err := row.Scan(
		&game.Id,
		&game.Name,
		&game.Description,
		&game.CreatedAt,
		&game.OwnerId,
		&game.ShareId,
	); err {
	case sql.ErrNoRows:
		return game, ErrNoMatch
	default:
		return game, err
	}
}

func (db Database) GetGamePayload(gameId string) (*models.GamePayLoad, error) {
	game, err := db.GetGame(gameId);
	players, err := db.GetGamePlayers(gameId)

	fmt.Println(players)

	payload := models.GamePayLoad{}
	payload.Game = *game
	payload.Players = *players

	return &payload, err
}
