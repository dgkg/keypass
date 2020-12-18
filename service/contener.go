package service

import (
	"log"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/segmentio/kafka-go"

	"github.com/dgkg/keypass/cache"
	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/middleware"
	"github.com/dgkg/keypass/model"
)

// ServiceContener reprensent all services around Contener.
type ServiceContener struct {
	DB    db.DB
	Cache cache.CacheDB
	Kw    *kafka.Writer
}

// @Description get a Contener by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "Some ID"
// @Success 200 {object} model.Contener "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /conteners/{uuid} [get]
func (su *ServiceContener) GetContener(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/conteners bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.GetContener(id.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
	}
	ctx.JSON(http.StatusOK, u)
}

// @Description create a Contener from the payload.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Contener body model.Contener true "Add a Contener"
// @Success 200 {object} model.Contener
// @Failure 400 {string} string nil
// @Router /conteners [post]
func (su *ServiceContener) CreateContener(ctx *gin.Context) {

	var u model.Contener
	err := ctx.BindJSON(&u)
	if err != nil {
		log.Println("/conteners bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	cession, ok := ctx.Get("cession")
	if !ok {
		return
	}
	t := cession.(*middleware.JWTClaims)
	u.UserID = t.UserUUID

	errs := u.ValidatePayload()
	if len(errs) != 0 {
		log.Println("/conteners bad request", errs)
		ctx.JSON(http.StatusBadRequest, errs)
		return
	}

	u2, _ := su.DB.CreateContener(&u)

	ctx.JSON(http.StatusOK, u2)
}

// @Description update a Contener from the payload.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Contener body model.Contener true "Add a Contener"
// @Success 200 {object} model.Contener
// @Failure 400 {string} string nil
// @Router /conteners [patch] [put]
func (su *ServiceContener) UpdateContener(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/conteners bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	var payload model.Payloadpatch
	payload.Data = make(map[string]interface{})
	err = ctx.BindJSON(&payload.Data)
	if err != nil {
		log.Println("/conteners bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.UpdateContener(id.String(), &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// DeleteContener is deleing a Contener form it's uuid.
// @Description delete a Contener by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "Some ID"
// @Success 200 {object} model.Contener "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /conteners/{uuid} [delete]
func (su *ServiceContener) DeleteContener(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/conteners bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = su.DB.DeleteContener(id.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// @Description get a Contener by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Contener "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /conteners [get]
func (su *ServiceContener) GetAllContener(ctx *gin.Context) {
	us, err := su.DB.GetAllContener()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, us)
}
