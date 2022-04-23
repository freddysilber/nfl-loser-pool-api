package db

import (
	"database/sql"
	"fmt"

	"github.com/freddysilber/nfl-loser-pool-api/models"
)

func (db Database) CreatePlayer(player *models.Player, id string) error {
	var gameId string
	var playerId string

	err := db.Conn.QueryRow(
		`
			INSERT INTO players (id, game_id, player_id)
			VALUES ($1, $2, $3)
			RETURNING id, game_id, player_id
		`,
		id,
		player.GameId,
		player.PlayerId,
	).Scan(
		&id,
		&gameId,
		&playerId,
	)

	if err != nil {
		return err
	}

	return nil
}

// Gets all the users from the player records for a given game
func (db Database) GetGamePlayers(gameId string) (*models.UserList, error) {
	list := &models.UserList{}
	users, err := db.Conn.Query(
		`
			SELECT
				u.id,
				name,
				username
			FROM users u
			JOIN players p ON p.player_id = u.id
			WHERE p.game_id = $1
		`,
		gameId,
	)

	if err != nil {
		return list, err
	}

	for users.Next() {
		var user models.User
		err := users.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
		)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}

	return list, nil
}

// Check the db to see if the new player already exists
func (db Database) GetExistingPlayer(player *models.Player) bool {
	var gameId string
	var playerId string

	row := db.Conn.QueryRow(
		`SELECT p.game_id, p.player_id FROM players p WHERE p.game_id = $1 AND p.player_id = $2`,
		player.GameId,
		player.PlayerId,
	)

	fmt.Println("row ")

	switch err := row.Scan(&gameId, &playerId); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		return false
	}
}
