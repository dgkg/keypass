package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/dgkg/keypass/db/mysql"
	_ "github.com/dgkg/keypass/docs"
	"github.com/dgkg/keypass/service"
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

// @host localhost:9090
// @BasePath /v1
func main() {

	r := gin.Default()

	var su service.ServiceUser
	// su.DB = moke.New()
	su.DB = mysql.New()

	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := r.Group("/v1")
	v1.GET("/users/:uuid", su.GetUser)
	v1.GET("/users", su.GetAllUser)
	v1.GET("/users/:uuid/*action", su.SetUserAction)
	v1.PATCH("/users/:uuid", su.UpdateUser)
	v1.PUT("/users/:uuid", su.UpdateUser)
	v1.POST("/users", su.CreateUser)
	v1.DELETE("/users/:uuid", su.DeleteUser)

	r.Run(":9090")
}
