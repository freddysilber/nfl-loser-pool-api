package handler

import (
	"log"
	"net/http"

	"github.com/freddysilber/nfl-loser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func players(router chi.Router) {
	router.Get("/", getAllPlayers)
	router.Post("/", createPlayer)
}

func getAllPlayers(w http.ResponseWriter, r *http.Request) {
	log.Println("get all players")
	userId := r.URL.Query().Get("player") // Get the '?player=x' query param value from the request
	log.Println("User Id --> ", userId)
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	_, err := ValidateSession(w, r)
	if err != nil {
		render.Render(w, r, UnAuthorized)
	}

	id, err := gonanoid.New()
	if err != nil {
		return
	}

	player := &models.Player{}

	if err := render.Bind(r, player); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if dbInstance.GetExistingPlayer(player) {
		render.Render(w, r, &ErrorResponse{StatusCode: 405, Message: "There is already a player for this user and game"})
		return
	}

	if err := dbInstance.CreatePlayer(player, id); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, player); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
