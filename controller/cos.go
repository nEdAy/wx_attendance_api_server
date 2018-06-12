package controller

import (
	"github.com/nEdAy/wx_attendance_api_server/internal/wx_cos_auth"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Binding from getAuthorizationModel JSON
type getAuthorizationModel struct {
	Method   string `json:"method" binding:"required"`
	Pathname string `json:"pathname" binding:"required"`
}

// GetAuthorization 获取鉴权签名
func GetAuthorization(c *gin.Context) {
	getAuthorizationModel := new(getAuthorizationModel)
	if err := c.ShouldBindWith(&getAuthorizationModel, binding.JSON); err != nil {
		renderJSONWithError(c, err.Error())
		return
	}
	authData, err := wx_cos_auth.AuthorizationTransport(getAuthorizationModel.Method, getAuthorizationModel.Pathname)
	if err != nil {
		renderJSONWithError(c, err.Error())
		return
	}
	if authData == "" {
		renderJSONWithError(c, "未成功生成鉴权签名")
		return
	}
	c.JSON(http.StatusOK, authData)
}
