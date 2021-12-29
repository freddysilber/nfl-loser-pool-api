package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/freddysilber/nfl-looser-pool-api/db"
	"github.com/freddysilber/nfl-looser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type userIdKeyString string

// Claims struct for jwt token contents
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// cookie key
const sessionToken = "session-token"

var userIDKey userIdKeyString = "userID"
// TODO: use a secure key mounted during deployment
var jwtKey = []byte("ja93jalkdf092jlkadfh02h3lkdfiu0293lakndf0923haf93ja1h")

func users(router chi.Router) {
	router.Get("/", getAllUsers)
	router.Post("/signup", signUp)
	router.Post("/login", logIn)
	router.Delete("/logout", logout)
	router.Route("/{userId}", func(router chi.Router) {
		router.Use(UserContext)
		router.Get("/", getUser)
		router.Delete("/", deleteUser)
	})
}

func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		if userId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("user ID is required")))
			return
		}
		id, err := strconv.Atoi(userId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid user ID")))
		}
		ctx := context.WithValue(r.Context(), userIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func signUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	
	if err := dbInstance.SignUp(user); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	// create jwt token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// return err
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	// update cookie
	cookie := http.Cookie{
		Name:    sessionToken,
		Value:   tokenString,
		Expires: expirationTime,
	}
	http.SetCookie(w, &cookie)

	// return user info in response, such as roles
	user.Password = "" // sanitize

	render.Render(w, r, user)
}

func logIn(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	row, err := dbInstance.GetUserByUsernameAndPassword(user)

	if err != nil {
		log.Println("ROW", row)
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

// // logs out user by invalidating session token
// func handleMethodDelete(ctx *fasthttp.RequestCtx) error {
// var c fasthttp.Cookie
// c.SetKey(sessionToken)
// c.SetValue("")
// c.SetExpire(time.Now())
// ctx.Response.Header.SetCookie(&c)
// ctx.SetStatusCode(fasthttp.StatusOK)
// return nil
// }
func logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout")
	cookie := http.Cookie{
		Name:    sessionToken,
		Value:   "",
		Expires: time.Now(),
	}
	http.SetCookie(w, &cookie)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIDKey).(int)
	user, err := dbInstance.GetUserById(userId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIDKey).(int)
	err := dbInstance.DeleteUser(userId)
	
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dbInstance.GetAllUsers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
