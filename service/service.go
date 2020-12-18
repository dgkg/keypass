package service

import (
	"github.com/dgkg/keypass/cache"
	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/middleware"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func New(r *gin.Engine, db db.DB, cache cache.CacheDB, kw *kafka.Writer) {
	var su ServiceUser
	su.DB = db
	su.Cache = cache
	su.Kw = kw
	v1 := r.Group("/v1")
	v1.POST("/login", su.LoginUser)
	v1.GET("/users/:uuid", su.GetUser)
	v1.GET("/users", middleware.NewCacheMiddleware(cache), su.GetAllUser)
	v1.PATCH("/users/:uuid", su.UpdateUser)
	v1.PUT("/users/:uuid", su.UpdateUser)
	v1.POST("/users", su.CreateUser)
	v1.DELETE("/users/:uuid", su.DeleteUser)

	var sc ServiceCard
	sc.DB = db
	sc.Cache = cache
	sc.Kw = kw
	card := v1.Group("/cards").Use(middleware.NewJWTMiddleware())
	card.GET("/:uuid", sc.GetCard)
	card.GET("", sc.GetAllCard)
	card.PATCH("/:uuid", sc.UpdateCard)
	card.PUT("/:uuid", sc.UpdateCard)
	card.POST("", sc.CreateCard)
	card.DELETE("/:uuid", sc.DeleteCard)

	var ser ServiceContener
	ser.DB = db
	ser.Cache = cache
	ser.Kw = kw
	contener := v1.Group("/conteners").Use(middleware.NewJWTMiddleware())
	contener.GET("/:uuid", ser.GetContener)
	contener.GET("", ser.GetAllContener)
	contener.PATCH("/:uuid", ser.UpdateContener)
	contener.PUT("/:uuid", ser.UpdateContener)
	contener.POST("", ser.CreateContener)
	contener.DELETE("/:uuid", ser.DeleteContener)

}
