package service

import (
	"log"
	"net/http"

	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type ServiceUser struct {
	DB db.DB
}

func (su *ServiceUser) SetUserAction(c *gin.Context) {
	id := c.Param("uuid")
	action := c.Param("action")
	message := id + " is " + action
	c.String(http.StatusOK, message)
}

// @Description get a User by ID
// @Accept json
// @Produce json
// @Param uuid path string true "Some ID"
// @Success 200 {object} model.User "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /users/{uuid} [get]
func (su *ServiceUser) GetUser(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.GetUser(id.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
	}
	ctx.JSON(http.StatusOK, u)
}

// @Description create a User from the payload.
// @Accept json
// @Produce json
// @Param user body model.User true "Add a User"
// @Success 200 {object} model.User
// @Failure 400 {string} string nil
// @Router /users [post]
func (su *ServiceUser) CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.BindJSON(&u)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	errs := u.ValidatePayload()
	if len(errs) != 0 {
		log.Println("/users bad request", errs)
		ctx.JSON(http.StatusBadRequest, errs)
		return
	}

	u2, _ := su.DB.CreateUser(&u)

	ctx.JSON(http.StatusOK, u2)
}

// @Description update a User from the payload.
// @Accept json
// @Produce json
// @Param user body model.User true "Add a User"
// @Success 200 {object} model.User
// @Failure 400 {string} string nil
// @Router /users [patch] [put]
func (su *ServiceUser) UpdateUser(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	var payload model.Payloadpatch
	payload.Data = make(map[string]interface{})
	err = ctx.BindJSON(&payload.Data)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.UpdateUser(id.String(), &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// DeleteUser is deleing a user form it's uuid.
// @Description delete a User by ID
// @Accept json
// @Produce json
// @Param uuid path string true "Some ID"
// @Success 200 {object} model.User "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /users/{uuid} [delete]
func (su *ServiceUser) DeleteUser(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.DeleteUser(id.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// @Description get a User by ID
// @Accept json
// @Produce json
// @Success 200 {object} []model.User "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /users [get]
func (su *ServiceUser) GetAllUser(ctx *gin.Context) {
	us, err := su.DB.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, us)
}
