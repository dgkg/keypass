package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func main() {

	UsersDB = make(map[string]*User)

	router := fiber.New()

	// This handler will match /user/john but will not match /user/ or /user
	router.Get("/users/:uuid", func(ctx *fiber.Ctx) error {

		id, err := uuid.FromString(ctx.Params("uuid"))
		if err != nil {
			log.Println("/users bad request", err.Error())
			return ctx.Status(http.StatusBadRequest).JSON(nil)
		}

		u, ok := UsersDB[id.String()]
		if !ok {
			return ctx.Status(http.StatusNotFound).JSON(nil)
		}
		return ctx.Status(http.StatusOK).JSON(u)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.Get("/users/:uuid/*action", func(c *fiber.Ctx) error {
		id := c.Params("uuid")
		action := c.Params("action")
		message := id + " is " + action
		return c.Status(http.StatusOK).SendString(message)
	})

	// For each matched request Context will hold the route definition
	router.Post("/users", func(ctx *fiber.Ctx) error {
		var u User
		err := ctx.BodyParser(&u)
		if err != nil {
			log.Println("/users bad request", err.Error())
			return ctx.Status(http.StatusBadRequest).JSON(nil)
		}
		u2, err := NewUser(u.FirstName, u.LastName, u.Email, u.Password)
		if err != nil {
			log.Println("/users create user", err.Error())
			return ctx.Status(http.StatusInternalServerError).JSON(nil)
		}
		UsersDB[u2.ID] = u2
		return ctx.Status(http.StatusOK).JSON(u2)
	})

	router.Listen(":9090")
}

// UsersDB is a moke for DB.
var UsersDB map[string]*User

// User represent a single customer.
type User struct {
	ID           string    `json:"uuid"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreationDate time.Time `json:"creation_date"`
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
