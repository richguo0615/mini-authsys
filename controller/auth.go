package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/richguo0615/mini-authsys/constant"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User ...
// Custom object which can be stored in the claims
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GeneratorJWT(userId int32, userName string) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"username": userName,
		"exp": time.Now().Add(time.Hour * 2).Unix(), //過期時間
	})
	return token.SignedString([]byte(constant.JwtSecretKey))
}

func ParseJwtClaims(parseToken *jwt.Token) (claims jwt.MapClaims) {
	claims = parseToken.Claims.(jwt.MapClaims)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(hashPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}
