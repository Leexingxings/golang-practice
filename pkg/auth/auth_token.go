package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type BlogClaims struct {
	uid uint64 `json:"uid,string"`
}

func BuildJWT() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Unix(),
	})

	hmacSampleSecret := []byte("my_secret")

	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)
}
