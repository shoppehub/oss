package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoppehub/conf"
	oss2 "github.com/shoppehub/oss"
	_ "github.com/shoppehub/oss/cmd/docs"
	"github.com/sirupsen/logrus"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host spiderapi.dev.chemball.com:4000
// @BasePath /
func main() {
	// r := oss.New()

	// // 使用默认中间件（logger和recovery）
	// url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//oss.Start(r)

	conf.Init("")
	port := conf.GetInt("port")
	Port := 4000
	if port != 0 {
		Port = int(port)
	}
	logrus.Info("start server on " + fmt.Sprint(Port))
	r := gin.Default()
	oss2.Route(r)
	r.Run("0.0.0.0:" + fmt.Sprint(Port))
}
