package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/freddysilber/nfl-loser-pool-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func session(router chi.Router) {
	router.Get("/", getSession)
	router.Delete("/", deleteSession)
}

func getSession(w http.ResponseWriter, r *http.Request) {
	user, err := ValidateSession(w, r)
	if err != nil {
		return
	}

	// b, err := json.Marshal(user)
	_, err = json.Marshal(user)
	if err != nil {
		return
	}
	render.Render(w, r, user)
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    sessionToken,
		Value:   "",
		Expires: time.Now(),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
}

// ValidateSession and check user has at least one of the roles. returns WebUserObject object if session is valid
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
	fmt.Println("")
	fmt.Println("--------- Claims ---------")
	fmt.Println("Claims --> ", claims)
	fmt.Println("Expires At --> ", claims.ExpiresAt)
	fmt.Println("Time Now --> ", time.Now())
	fmt.Println("Time Now Unix --> ", time.Now().Unix())
	fmt.Println(claims.ExpiresAt < time.Now().Unix())
	if claims.ExpiresAt < time.Now().Unix() {
		return user, errors.New("session expired")
	} else {
		// I guess we should update the expiration time...
		/*
			The goal here is to keep the user logged into the app as long as they are active and making requests
			Everytime we Validate them (they should be authed and logged in) we should push their expiration time back
		*/
		// claims.ExpiresAt = time.Now().Add(60 * time.Minute).Unix()
		fmt.Println("Claims: Expiration time should have been changed...?", claims.ExpiresAt)
		fmt.Println("CLAIMS EXPIRATION SHOULD BE UPDATED (If it's not, we got work to do...)")
	}
	fmt.Println("--------- Claims ---------")
	fmt.Println("")
	// Todo: this should be fixed
	user.Username = claims.Username
	user.Id = claims.Id
	user.Password = claims.Password
	user, err = dbInstance.GetUserByIdUsernameAndPassword(user)
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
	tokenString, err := r.Cookie(sessionToken)
	if err != nil {
		return nil, err
	}
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
