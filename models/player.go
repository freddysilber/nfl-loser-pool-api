package models

import "net/http"

type Player struct {
	Id       string `json:"id"`
	GameId   string `json:"gameId"`
	PlayerId string `json:"playerId"`
	// ?
	// Realted User from 'player_id'
	// User User `json:"user"`
}

type PlayerList struct {
	Players []Player `json:"players"`
}

func (g *Player) Bind(r *http.Request) error {
	// if g.Name == "" {
	// 	return fmt.Errorf("name is a required field")
	// }
	return nil
}

func (*PlayerList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Player) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
