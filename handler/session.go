package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/freddysilber/nfl-looser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func session(router chi.Router) {
	router.Get("/", getSession)
	router.Post("/", makeSession)
	router.Delete("/", deleteSession)
}

func getSession(w http.ResponseWriter, r *http.Request) {
	user, err := ValidateSession(w, r)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	b, err := json.Marshal(user)
	log.Println("b -->", b)
	if err != nil {
		return
	}
	render.Render(w, r, user)
}

func makeSession(w http.ResponseWriter, r *http.Request) {

}

func deleteSession(w http.ResponseWriter, r *http.Request) {

}

// ValidateSession and check user has atleast one of the roles. returns WebUserObject object iff session is valid
func ValidateSession(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	user := &models.User{}
	token, err := verifyToken(r)
	if err != nil {
		return user, err
	}
	b, err := json.Marshal(token.Claims)
	if err != nil {
		return user, err
	}
	var claims Claims
	err = json.Unmarshal(b, &claims)
	if err != nil {
		return user, err
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return user, errors.New("session expired")
	}
	user.Username = claims.Username
	user, err = dbInstance.GetUserByUsernameAndPassword(user)
	if err != nil {
		return user, err
	}
	return user, nil
	// if len(roles) == 0 {
	// 	return user, nil
	// }
	// for _, requiredRole := range roles {
	// 	for _, myRole := range user.Roles {
	// 		if requiredRole == myRole {
	// 			return user, nil
	// 		}
	// 	}
	// }
	// return nil, errors.New("user doesn't have role")
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, _ := r.Cookie(sessionToken)
	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}