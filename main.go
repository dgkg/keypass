package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func main() {

	UsersDB = make(map[string]*User)

	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/users/:uuid", func(ctx *gin.Context) {

		id, err := uuid.FromString(ctx.Param("uuid"))
		if err != nil {
			log.Println("/users bad request", err.Error())
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		u, ok := UsersDB[id.String()]
		if !ok {
			ctx.JSON(http.StatusNotFound, nil)
		}
		ctx.JSON(http.StatusOK, u)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/users/:uuid/*action", func(c *gin.Context) {
		id := c.Param("uuid")
		action := c.Param("action")
		message := id + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/users", func(ctx *gin.Context) {
		var u User
		err := ctx.BindJSON(&u)
		if err != nil {
			log.Println("/users bad request", err.Error())
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		UsersDB[u.ID] = &u
		ctx.JSON(http.StatusOK, u)
	})

	router.Run(":9090")
}

// UsersDB is a moke for DB.
var UsersDB map[string]*User

// User represent a single customer.
type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	CreationDate time.Time
}

func NewUser(fn, ln, email, pass string) (*User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &User{
		ID:           id.String(),
		FirstName:    fn,
		LastName:     ln,
		Email:        email,
		Password:     pass,
		CreationDate: time.Now(),
	}, nil
}
