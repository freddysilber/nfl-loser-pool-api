package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/freddysilber/nfl-loser-pool-api/db"
	"github.com/freddysilber/nfl-loser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type gameIdKeyString string

var gameIdKey gameIdKeyString = "gameId"

func games(router chi.Router) {
	router.Get("/", getAllGames)
	router.Post("/", createGame)
	router.Route("/{gameId}", func(router chi.Router) {
		router.Use(GameContext)
		router.Delete("/", deleteGame)
		router.Get("/players", getGamePlayers)
	})
}

func GameContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gameId := chi.URLParam(r, "gameId")
		if gameId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("game ID is required")))
			return
		}
		id, err := strconv.Atoi(gameId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid game ID")))
		}
		ctx := context.WithValue(r.Context(), gameIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllGames(w http.ResponseWriter, r *http.Request) {
	// TODO: abstract this block
	_, err := ValidateSession(w, r)
	if err != nil {
		render.Render(w, r, UnAuthorized)
	}
	games, err := dbInstance.GetAllGames()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, games); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createGame(w http.ResponseWriter, r *http.Request) {
	// TODO: abstract this block
	_, err := ValidateSession(w, r)
	if err != nil {
		render.Render(w, r, UnAuthorized)
	}
	// Create a random game id / uuid
	id, err := gonanoid.New()
	if err != nil {
		return
	}
	game := &models.Game{}
	if err := render.Bind(r, game); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddGame(game, id); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, game); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	gameId := r.Context().Value(gameIdKey).(int)
	err := dbInstance.DeleteGame(gameId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func getGamePlayers(w http.ResponseWriter, r *http.Request) {
	// TODO this pops up in muliple places, so we should find a solution to remove this block
	_, err := ValidateSession(w, r)
	if err != nil {
		render.Render(w, r, UnAuthorized)
	}

	gameId := r.Context().Value(gameIdKey).(int)
	players, err := dbInstance.GetGamePlayers(gameId)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, players); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
