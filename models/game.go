package models

import (
	"fmt"
	"net/http"
)

type Game struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     int    `json:"ownerId"`
	CreatedAt   string `json:"created_at"`
	ShareId     string `json:"share_id"`
}

type GameList struct {
	Games []Game `json:"games"`
}

func (g *Game) Bind(r *http.Request) error {
	if g.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*GameList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Game) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
