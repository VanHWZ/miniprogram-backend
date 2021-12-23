package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	cstZone = time.FixedZone("CST", 8*3600)
	lease = time.Hour * 8
	jwtSecret = []byte("J&VEvent")
)

func ValidateToken(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err == nil{
		return true
	} else{
		return false
	}
}

func NewToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().In(cstZone).Add(lease).Unix(),
	})
	ts, _ := token.SignedString(jwtSecret)
	return ts
}