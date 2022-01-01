package handler

import (
	"log"
	"net/http"

	"github.com/freddysilber/nfl-loser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func games(router chi.Router) {
	router.Get("/", getAllGames)
	router.Post("/", createGame)
}

func getAllGames(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllGames")
}

func createGame(w http.ResponseWriter, r *http.Request) {
	_, err := ValidateSession(w, r)
	if err != nil {
		render.Render(w, r, UnAuthorized)
	}

	game := &models.Game{}
	if err := render.Bind(r, game); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddGame(game); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, game); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}