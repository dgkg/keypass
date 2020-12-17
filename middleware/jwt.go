package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("AllYourBase")

func JWTKey(value string) {
	mySigningKey = []byte(value)
}

type JWTClaims struct {
	UserUUID string `json:"uuid"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func NewJWT(uuid, userName string) (string, error) {

	// Create the Claims
	claims := JWTClaims{
		UserUUID: uuid,
		UserName: userName,
	}
	claims.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
	claims.Issuer = "keypass"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ValidateJWT(tokenValue string) (*JWTClaims, error) {

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(JWTClaims); ok && token.Valid {
		fmt.Println(claims.UserUUID, claims.UserName)
		return &claims, nil
	}
	return nil, err
}
