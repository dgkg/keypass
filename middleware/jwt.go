package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	// claims.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
	claims.ExpiresAt = time.Now().Unix()
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

func NewJWTMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		value := ctx.Request.Header.Get("Authorization")
		if len(value) < 190 {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("JWT no value"))
			return
		}
		valueToken := strings.Split(value, " ")
		if len(valueToken) != 2 {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("JWT parsing"))
			return
		}
		_, err := ValidateJWT(valueToken[1])
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
	}
}
