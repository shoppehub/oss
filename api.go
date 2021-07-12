package oss

import (
	"net/http"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/shoppehub/conf"
	"github.com/shoppehub/oss/service"
	"github.com/sirupsen/logrus"
)

// @Summary aliyun oss upload policy
// @Description
// @Accept  json
// @Produce  json
// @Param upload_dir query string false "上传的目录，默认是 upload"
// @Success 200 {object} service.PolicyToken	"ok"
// @Router /api/oss/getconfig [get]
func getOssConfig(ctx *gin.Context) {

	service.Init()

	upload_dir := ctx.Query("upload_dir")

	response := service.Get_policy_token(upload_dir)

	ctx.JSON(http.StatusOK, response)

}

var client *oss.Client

func initOSS() {
	c, err := oss.New(conf.GetString("oss.endpoint"), conf.GetString("oss.accessKeyId"), conf.GetString("oss.accessKeySecret"))
	if err != nil {
		logrus.Error(err)
		return
	}
	client = c
}

func delOss(ctx *gin.Context) {

	key := ctx.Query("key")
	bucketName := conf.GetString("oss.bucket")

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		logrus.Error(err, key)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})

}
