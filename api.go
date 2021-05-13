package oss

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoppehub/oss/service"
)

// @Summary aliyun oss upload policy
// @Description
// @Accept  json
// @Produce  json
// @Param upload_dir query string false "上传的目录，默认是 upload"
// @Success 200 {object} service.PolicyToken	"ok"
// @Router /api/oss/getconfig [get]
func getOssConfig(ctx *gin.Context) {

	upload_dir := ctx.Query("upload_dir")

	response := service.Get_policy_token(upload_dir)

	ctx.JSON(http.StatusOK, response)

}
