package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/dgkg/keypass/docs"
	"github.com/dgkg/keypass/model"
)

// @title Swagger For Keypass API
// @version 1.0
// @description This is an API for creating hash in order to create keypasses.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	DB.users = make(map[string]*model.User)

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/users/:uuid", GetUser)
	r.GET("/users/:uuid/*action", SetUserAction)
	r.PATCH("/users/:uuid", UpdateUser)
	r.PUT("/users/:uuid", UpdateUser)

	r.POST("/users", CreateUser)

	r.Run(":9090")
}

func SetUserAction(c *gin.Context) {
	id := c.Param("uuid")
	action := c.Param("action")
	message := id + " is " + action
	c.String(http.StatusOK, message)
}

// @Description get a User by ID
// @Accept json
// @Produce json
// @Param uuid path string true "Some ID"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /users/{some_id} [get]
func GetUser(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, ok := DB.users[id.String()]
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
	DB.users[u2.ID] = u2
	ctx.JSON(http.StatusOK, u2)
}

func UpdateUser(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	var payload Payloadpatch
	payload.data = make(map[string]interface{})
	err = ctx.BindJSON(&payload.data)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, ok := DB.users[id.String()]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	u.FirstName = payload.toString("first_name")
	u.LastName = payload.toString("last_name")
	u.Email = payload.toString("email")

	if len(payload.errs) != 0 {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

type Payloadpatch struct {
	data map[string]interface{}
	errs []error
}

func (p *Payloadpatch) toString(fieldName string) string {
	val, ok := p.data[fieldName]
	if !ok {
		p.errs = append(p.errs, errors.New("no values for:"+fieldName))
		return ""
	}
	newval, ok := val.(string)
	if !ok {
		p.errs = append(p.errs, errors.New("cast not possible into string for:"+fieldName))
		return ""
	}
	return newval
}

type mokeDB struct {
	users map[string]*model.User
}

// DB is a moke for DB.
var DB mokeDB
