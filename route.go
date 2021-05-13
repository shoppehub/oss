package oss

import "github.com/gin-gonic/gin"

func route(r *gin.Engine) {
	r.GET("/api/oss/getconfig", getOssConfig)

	r.POST("/api/oss/callback", getOssConfig)

}
