package main

import (
	"log"
	"net/http"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"

	"github.com/dgkg/keypass/model"
)

func main() {

	UsersDB = make(map[string]*model.User)

	router := gin.Default()

	router.GET("/users/:uuid", GetUser)
	router.GET("/users/:uuid/*action", SetUserAction)
	router.POST("/users", CreateUser)

	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":9090")
}

func SetUserAction(c *gin.Context) {
	id := c.Param("uuid")
	action := c.Param("action")
	message := id + " is " + action
	c.String(http.StatusOK, message)
}

func GetUser(ctx *gin.Context) {

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
}

func CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.BindJSON(&u)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	u2, err := model.NewUser(u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		log.Println("/users create user", err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	UsersDB[u2.ID] = u2
	ctx.JSON(http.StatusOK, u2)
}

// UsersDB is a moke for DB.
var UsersDB map[string]*model.User
