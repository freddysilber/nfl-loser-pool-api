package db

import (
	"encoding/json"
	"log"

	"github.com/freddysilber/nfl-loser-pool-api/models"
)

// Gets all the 'Player' records from a game id
func (db Database) GetGamePlayers(gameId string) (* models.PlayerList, error) {
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