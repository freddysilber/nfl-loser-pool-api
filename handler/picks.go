package handler

import (
	"net/http"

	"github.com/freddysilber/nfl-loser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func picks(router chi.Router) {
	// router.Get("/", getAllPicks)
	router.Post("/", createPick)
	// router.Route("/{}")
}

// func getAllPicks(w http.ResponseWriter, r *http.Request) {
// 	_, err := ValidateSession(w, r)
// 	if err != nil {
// 		render.Render(w, r, UnAuthorized)
// 	}

// 	picks, err := dbInstance.GetAllPicks()
// }

func createPick(w http.ResponseWriter, r *http.Request) {
	_, err := ValidateSession(w, r)
	if err != nil {
		render.Render(w, r, UnAuthorized)
	}

	id, err := gonanoid.New()
	if err != nil {
		return
	}

	pick := &models.Pick{}

	if err := dbInstance.CreatePick(pick, id); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}