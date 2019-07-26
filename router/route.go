package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/richguo0615/mini-authsys/constant"
	"github.com/richguo0615/mini-authsys/controller/defalt"
	"github.com/richguo0615/mini-authsys/controller/user"
	"github.com/richguo0615/mini-authsys/helper"
	"net/http"
)

func InitRoute() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", user.SignUp).Methods("POST")
	router.HandleFunc("/auth", user.Auth).Methods("POST")
	router.HandleFunc("/userInfo", validateTokenMiddleware(user.GetUserInfo)).Methods("GET")
	router.HandleFunc("/trans", validateTokenMiddleware(user.Trans)).Methods("POST")
	router.HandleFunc("/", defalt.Index).Methods("GET")

	err := http.ListenAndServe(":8082", router)
	if err != nil {
		panic(err)
	}
}

// ErrorMsg ...
// Custom error object
type ErrorMsg struct {
	Message string `json:"message"`
}

//驗證jwt
func validateTokenMiddleware(next func(w http.ResponseWriter, r *http.Request, parseToken *jwt.Token)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tokenStr := req.Header.Get("authorization")
		if len(tokenStr) == 0 {
			helper.ResponseWithJson(w, http.StatusUnauthorized, "not authorized")
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					helper.ResponseWithJson(w, http.StatusUnauthorized, "not authorized")
					return nil, fmt.Errorf("not authorized")
				}
				return []byte(constant.JwtSecretKey), nil
			})
			if !token.Valid {
				helper.ResponseWithJson(w, http.StatusUnauthorized, "not authorized")
			} else {
				next(w, req, token)
			}
		}
	})
}
