package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/db/mysql"
	"github.com/dgkg/keypass/db/sqlite"
	_ "github.com/dgkg/keypass/docs"
	"github.com/dgkg/keypass/service"
)

type Config struct {
	MySQL string
	Mode  string
	Port  string
}

var conf Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	conf.MySQL = viper.GetString("mysql.dsn")
	conf.Mode = viper.GetString("mode")
	conf.Port = viper.GetString("port")
}

// @title Swagger For Keypass API
// @version 1.0
// @description This is an API for creating hash in order to create keypasses.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:" + conf.Port + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	fmt.Println("run app in mode:", conf.Mode)

	var db db.DB
	if conf.Mode == "local" {
		db = sqlite.New("local.db")
	} else {
		db = mysql.New(conf.MySQL)
	}

	service.New(r, db)

	r.Run(":" + conf.Port)
}
