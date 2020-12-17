package service

import (
	"github.com/dgkg/keypass/cache"
	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/middleware"
	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db db.DB, cache cache.CacheDB) {
	var su ServiceUser
	su.DB = db
	su.Cache = cache
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
	card := v1.Group("/cards").Use(middleware.NewJWTMiddleware())
	card.GET("/:uuid", sc.GetCard)
	card.GET("", sc.GetAllCard)
	card.PATCH("/:uuid", sc.UpdateCard)
	card.PUT("/:uuid", sc.UpdateCard)
	card.POST("", sc.CreateCard)
	card.DELETE("/:uuid", sc.DeleteCard)
}
