package db

import "github.com/freddysilber/nfl-loser-pool-api/models"

func (db Database) CreatePick(pick *models.Pick, id string) error {
	err := db.Conn.QueryRow(
		`
			INSERT INTO picks (id)
			VALUES ($1)
			RETURNING id	
		`,
		id,
	).Scan(
		&id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db Database) GetGamePicks(gameId string) (*models.PickList, error) {
	list := &models.PickList{}
	picks, err := db.Conn.Query(
		`
			SELECT *
			FROM picks
			WHERE game_id = $1
		`,
		gameId,
	)

	if err != nil {
		return list, err
	}

	for picks.Next() {
		var pick models.Pick
		err := picks.Scan(
			&pick.Id,
			&pick.GameId,
			&pick.PlayerId,
			&pick.TeamId,
			&pick.Week,
			&pick.BonusPoint,
		)
		if err != nil {
			return list, err
		}
		list.Picks = append(list.Picks, pick)
	}

	return list, nil
}
