package models

import "net/http"

type Pick struct {
	Id         string `json:"id"`
	GameId     string `json:"gameId"`
	PlayerId   string `json:"playerId"`
	TeamId     string `json:"teamId"`
	Week       int    `json:"week"`
	BonusPoint int    `json:"bonusPoint"`
}

type PickList struct {
	Picks []Pick `json:"picks"`
}

func (g *Pick) Bind(r *http.Request) error {
	return nil
}

func (*Pick) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*PickList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
