package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func session(router chi.Router) {
	router.Get("/", getSession)
	router.Post("/", makeSession)
	router.Delete("/", deleteSession)
}

func getSession(w http.ResponseWriter, r *http.Request) {
	log.Println("getSession")
	// user, err := ValidateSession(ctx)
	// if err != nil {
	// 	ctx.SetStatusCode(fasthttp.StatusUnauthorized)
	// 	return nil
	// }

	// // return user info in response, such as roles
	// b, err := json.Marshal(user)
	// if err != nil {
	// 	return err
	// }
	// ctx.SetBody([]byte(b))
	// ctx.SetStatusCode(fasthttp.StatusOK)
	// return nil
}

func makeSession(w http.ResponseWriter, r *http.Request) {

}

func deleteSession(w http.ResponseWriter, r *http.Request) {

}

// ValidateSession and check user has atleast one of the roles. returns WebUserObject object iff session is valid
// func ValidateSession(ctx *fasthttp.RequestCtx, roles ...string) (*user.WebUserObject, error) {
// 	token, err := verifyToken(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	b, err := json.Marshal(token.Claims)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var claims Claims
// 	err = json.Unmarshal(b, &claims)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if claims.ExpiresAt < time.Now().Unix() {
// 		return nil, errors.New("session expired")
// 	}
// 	user, err := user.GetWebUserObject(claims.Username, ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(roles) == 0 {
// 		return user, nil
// 	}
// 	for _, requiredRole := range roles {
// 		for _, myRole := range user.Roles {
// 			if requiredRole == myRole {
// 				return user, nil
// 			}
// 		}
// 	}
// 	return nil, errors.New("user doesn't have role")
// }