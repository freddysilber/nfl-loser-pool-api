package models

import (
	"fmt"
	"net/http"
)

type Game struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     string `json:"ownerId"`
	CreatedAt   string `json:"createdAt"`
	ShareId     string `json:"shareId"`
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
