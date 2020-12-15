package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/dgkg/keypass/db/moke"
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

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	r := gin.Default()

	var su service.ServiceUser
	su.DB = moke.New()

	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/users/:uuid", su.GetUser)
	r.GET("/users/:uuid/*action", su.SetUserAction)
	r.PATCH("/users/:uuid", su.UpdateUser)
	r.PUT("/users/:uuid", su.UpdateUser)
	r.POST("/users", su.CreateUser)

	r.Run(":9090")
}
