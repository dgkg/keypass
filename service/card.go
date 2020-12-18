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
	"github.com/dgkg/keypass/model"
)

// ServiceCard reprensent all services around card.
type ServiceCard struct {
	DB    db.DB
	Cache cache.CacheDB
	Kw    *kafka.Writer
}

// @Description get a Card by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "Some ID"
// @Success 200 {object} model.Card "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /cards/{uuid} [get]
func (su *ServiceCard) GetCard(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/cards bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.GetCard(id.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
	}
	ctx.JSON(http.StatusOK, u)
}

// @Description create a Card from the payload.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Card body model.Card true "Add a Card"
// @Success 200 {object} model.Card
// @Failure 400 {string} string nil
// @Router /cards [post]
func (su *ServiceCard) CreateCard(ctx *gin.Context) {
	var u model.Card
	err := ctx.BindJSON(&u)
	if err != nil {
		log.Println("/cards bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	errs := u.ValidatePayload()
	if len(errs) != 0 {
		log.Println("/cards bad request", errs)
		ctx.JSON(http.StatusBadRequest, errs)
		return
	}

	u2, _ := su.DB.CreateCard(&u)

	ctx.JSON(http.StatusOK, u2)
}

// @Description update a Card from the payload.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Card body model.Card true "Add a Card"
// @Success 200 {object} model.Card
// @Failure 400 {string} string nil
// @Router /cards [patch] [put]
func (su *ServiceCard) UpdateCard(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/cards bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	var payload model.Payloadpatch
	payload.Data = make(map[string]interface{})
	err = ctx.BindJSON(&payload.Data)
	if err != nil {
		log.Println("/cards bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.UpdateCard(id.String(), &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// DeleteCard is deleing a Card form it's uuid.
// @Description delete a Card by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "Some ID"
// @Success 200 {object} model.Card "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /cards/{uuid} [delete]
func (su *ServiceCard) DeleteCard(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("uuid"))
	if err != nil {
		log.Println("/cards bad request", err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	u, err := su.DB.DeleteCard(id.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// @Description get a Card by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Card "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /cards [get]
func (su *ServiceCard) GetAllCard(ctx *gin.Context) {
	us, err := su.DB.GetAllCard()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, us)
}
