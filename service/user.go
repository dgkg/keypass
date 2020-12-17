package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgkg/keypass/cache"
	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/middleware"
	"github.com/dgkg/keypass/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// ServiceUser reprensent all services around user.
type ServiceUser struct {
	DB    db.DB
	Cache cache.CacheDB
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

// @Description check if the login pass are correct and gives backe a JWT value
// @Accept json
// @Produce json
// @Param user body model.UserLogin true "Add a User"
// @Success 200 {object} string nil
// @Failure 400 {string} string nil
// @Router /login [post]
func (su *ServiceUser) LoginUser(ctx *gin.Context) {
	var payload model.UserLogin
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u2, err := su.DB.GetUserByEmail(payload.Login)
	if err != nil {
		log.Println("/users not found", err.Error())
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	if u2.Password != payload.Password {
		log.Println("/users not authorized")
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	jwtValue, err := middleware.NewJWT(u2.ID, u2.FirstName+" "+strings.ToUpper(u2.LastName))
	if err != nil {
		log.Println("/users not internal server error", err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
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

	su.Cache.Set(ctx, ctx.Request.RequestURI, us)

	ctx.JSON(http.StatusOK, us)
}
