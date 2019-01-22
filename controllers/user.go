package controllers

import (
	"GoWebAPI-JWT-Authorize/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 用户登录
// @Summary 登录
// @Tags 账号
// @Produce  json
// @Param userName query string true "userName"
// @Param userPwd query string true "userPwd"
// @Success 200 {string} json "{"Code":0,"Msg":"","Body":"abc"}"
// @Router /Login [get]
func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.ResultModel{
		Code: 0,
		Msg:  "",
		Body: "abc",
	})
}
