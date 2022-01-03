package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
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
	// TODO: User Id and Game Id need to be required!!
	log.Println("create player")
}
