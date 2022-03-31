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
	/*
		TODO: !! WE SERIOUSLY NEED TO FIND A WAY TO ABSTRACT THIS!!
		Take care of our user validation
	*/
	user, err := ValidateSession(w, r)
	if err != nil {
		render.Render(w, r, UnAuthorized)
	}

	// TODO: User Id and Game Id need to be required!!
	log.Println("create player, we are valid")
	log.Println("here is the valid user: ", user)
	log.Println(user.Id)
	log.Println(user.Name)
	log.Println(user.Username)
	log.Println(user.Password)

	// TODO: PULL THIS INTO A METHOD
	playerId, err := gonanoid.New()
	if err != nil {
		return
	}

	player := &models.Player{}

	if err := render.Bind(r, player); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.CreatePlayer(player, playerId); err != nil {
		log.Println(w)
		log.Println(r)
		log.Println("bro")
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, player); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
