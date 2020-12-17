package service

import (
	"github.com/dgkg/keypass/db"
	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db db.DB) {
	var su ServiceUser
	su.DB = db
	v1 := r.Group("/v1")
	v1.POST("/login", su.LoginUser)
	v1.GET("/users/:uuid", su.GetUser)
	v1.GET("/users", su.GetAllUser)
	v1.PATCH("/users/:uuid", su.UpdateUser)
	v1.PUT("/users/:uuid", su.UpdateUser)
	v1.POST("/users", su.CreateUser)
	v1.DELETE("/users/:uuid", su.DeleteUser)

	var sc ServiceCard
	sc.DB = db
	v1.GET("/cards/:uuid", sc.GetCard)
	v1.GET("/cards", sc.GetAllCard)
	v1.PATCH("/cards/:uuid", sc.UpdateCard)
	v1.PUT("/cards/:uuid", sc.UpdateCard)
	v1.POST("/cards", sc.CreateCard)
	v1.DELETE("/cards/:uuid", sc.DeleteCard)
}
