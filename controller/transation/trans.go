package transation

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/richguo0615/mini-authsys/controller"
	"github.com/richguo0615/mini-authsys/helper"
	"github.com/richguo0615/mini-authsys/kafka"
	"github.com/richguo0615/mini-authsys/model"
	"github.com/richguo0615/mini-authsys/model/db"
	"github.com/rs/xid"
	"net/http"
)

func TransIn(w http.ResponseWriter, r *http.Request, parseToken *jwt.Token) {

	fmt.Println("transIn!")

	defer r.Body.Close()

	var transaction db.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	claims := controller.ParseJwtClaims(parseToken)
	userId := claims["userId"].(float64)

	//送消息
	transaction.Serial = xid.New().String()
	transaction.OperType = model.OperTypeTransIn
	transaction.UserId = int32(userId)

	err := kafka.SendTransIn(transaction)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusExpectationFailed, model.Result{
			Code: model.MsgCode_MsgCode_FAIL,
			Msg:  fmt.Sprint(err),
		})
		return
	}

	helper.ResponseWithJson(w, http.StatusCreated, model.Result{
		Code: model.MsgCode_MsgCode_SUCCESS,
		Msg: "req success",
	})
}
