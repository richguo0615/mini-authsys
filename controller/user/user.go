package user

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/richguo0615/mini-authsys/conf"
	"github.com/richguo0615/mini-authsys/controller"
	"github.com/richguo0615/mini-authsys/helper"
	"github.com/richguo0615/mini-authsys/model"
	"github.com/richguo0615/mini-authsys/model/db"
	"net/http"
)

func SignUp(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("signup!")

	defer request.Body.Close()

	var user db.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		helper.ResponseWithJson(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	hashPassword, err := controller.HashPassword(user.Password)
	if err != nil {
		helper.ResponseWithJson(writer, http.StatusBadRequest, "hash password failed")
		return
	}
	user.Password = hashPassword

	if dbc := conf.DB.Create(&user); dbc.Error != nil {
		helper.ResponseWithJson(writer, http.StatusBadRequest, dbc.GetErrors())
		return
	}

	jwtStr, err := controller.GeneratorJWT(user.Name)
	if err != nil {
		helper.ResponseWithJson(writer, http.StatusBadRequest, err)
		return
	}

	helper.ResponseWithJson(writer, http.StatusCreated, model.UserRes{
		Result: model.Result{
			Code: model.MsgCode_MsgCode_SUCCESS,
		},
		Token: jwtStr,
	})
}

func Auth(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("auth login!")
	defer request.Body.Close()

	var userData db.User
	if err := json.NewDecoder(request.Body).Decode(&userData); err != nil {
		helper.ResponseWithJson(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user db.User
	conf.DB.Where("email = ?", userData.Email).First(&user)

	err := controller.CheckPassword(user.Password, userData.Password)
	if err != nil {
		helper.ResponseWithJson(writer, http.StatusBadRequest, "auth is failed.")
		return
	}

	jwtStr, err := controller.GeneratorJWT(user.Name)
	if err != nil {
		helper.ResponseWithJson(writer, http.StatusBadRequest, err)
	}

	helper.ResponseWithJson(writer, http.StatusCreated, model.UserRes{
		Result: model.Result{
			Code: model.MsgCode_MsgCode_SUCCESS,
		},
		Token: jwtStr,
	})
}

func GetUserInfo(writer http.ResponseWriter, request *http.Request, parseToken *jwt.Token) {
	fmt.Println("get user info!")
	defer request.Body.Close()

	var user db.User
	claims := controller.ParseJwtClaims(parseToken)
	conf.DB.Where("name = ?", claims["username"]).First(&user)

	user.Password = ""

	helper.ResponseWithJson(writer, http.StatusOK, user)
}

func Trans(w http.ResponseWriter, r *http.Request, parseToken *jwt.Token) {

	fmt.Println("trans!")

	defer r.Body.Close()

	var transaction db.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
}
