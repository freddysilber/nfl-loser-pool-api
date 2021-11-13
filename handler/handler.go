package handler

import (
	"net/http"

	"github.com/freddysilber/nfl-looser-pool-api/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		// MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	dbInstance = db

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	// Register model specific route handlers
	router.Route("/items", items)
	router.Route("/users", users)
	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// When there is a 400/ or there is a bad request?
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
