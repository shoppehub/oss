package oss

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	// 使用默认中间件（logger和recovery）
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	// api.Create(r)

	route(r)

	return r
}

func Start(r *gin.Engine) {
	// 启动服务，并默认监听4000端口
	port := "4000"
	log.Println("start server on " + port)

	r.Run(":" + port)
}
