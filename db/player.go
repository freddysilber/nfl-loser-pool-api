package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

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

// Gets all the 'Player' records from a game id
func (db Database) GetGamePlayers(gameId string) (*models.PlayerList, error) {
	list := &models.PlayerList{}
	players, err := db.Conn.Query(
		`
			SELECT
				p.id,
				p.game_id,
				p.player_id,
				u.id,
				u.name,
				u.username
			FROM players p
			INNER JOIN users u
			ON u.id = p.player_id
			WHERE p.game_id = $1
		`,
		gameId,
	)
	// players, err := db.Conn.Query(
	// 	`
	// 		SELECT *
	// 		FROM players
	// 		WHERE game_id = $1
	// 	`,
	// 	gameId,
	// )
	log.Println(list)
	log.Println(json.Marshal(list))

	if err != nil {
		return list, err
	}

	for players.Next() {
		var player models.Player
		var user models.User
		err := players.Scan(
			&player.Id,
			&player.GameId,
			&player.PlayerId,
			&user.Id,
			&user.Name,
			&user.Username,
		)
		if err != nil {
			return list, err
		}
		list.Players = append(list.Players, player)
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
