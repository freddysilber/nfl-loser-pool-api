package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/freddysilber/nfl-looser-pool-api/db"
	"github.com/freddysilber/nfl-looser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type userIdKeyString string

// Claims struct for jwt token contents
type Claims struct {
	Id int `json:"id"`// User Id
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// Credentials struct for demarshalling session post body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	
	if err := dbInstance.NewUser(user); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	// create jwt token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		Id: user.Id,
		Password: user.Password,
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
		Path: "/",
	}
	http.SetCookie(w, &cookie)

	// return user info in response, such as roles
	user.Password = "" // sanitize

	render.Render(w, r, user)
}

func logIn(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest);
		return
	}
	var providedPassword = user.Password
	
	user, err := dbInstance.GetUserByUsername(user)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	if user == nil {
		render.Render(w, r, UnAuthorized)
		return
	}

	err = verifyPassword(user.Password, []byte(providedPassword))
	if err != nil {
		render.Render(w, r, UnAuthorized)
		return
	}

	// create jwt token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		Id: user.Id,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	// update cookie
	cookie := http.Cookie{
		Name:    sessionToken,
		Value:   tokenString,
		Expires: expirationTime,
		Path: "/",
	}
	http.SetCookie(w, &cookie)
	render.Render(w, r, user)
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

// compaire plain-text password against a hashed-and-salted password
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func verifyPassword(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return err
	}
	return nil
}
