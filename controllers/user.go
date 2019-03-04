package controllers

import (
	"net/http"

	"github.com/zboyco/GoWebAPI-JWT-Authorize/models"

	"github.com/gin-gonic/gin"
)

// Login 用户登录
// @Summary 登录
// @Tags 账号
// @Produce  json
// @Param userName query string true "userName"
// @Param userPwd query string true "userPwd"
// @Success 200 object models.ResultModel
// @Router /Login [get]
func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.ResultModel{
		Code: 0,
		Msg:  "",
		Body: "abc",
	})
}
