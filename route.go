package oss

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {

	r.GET("/api/oss/getconfig", getOssConfig)

	r.POST("/api/oss/del", delOss)

	initOSS()

	imOss := r.Group("/api/im/oss")
	{
		imOss.GET("/getImOssConfig", getImOssConfig)
	}
	initImOSS()

	// r.POST("/api/oss/callback", getOssConfig)

}
