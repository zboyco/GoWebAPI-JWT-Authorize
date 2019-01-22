package controllers

import (
	"GoWebAPI-JWT-Authorize/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccountInfo 获取账户信息
// @Summary 获取账户信息
// @Tags 账户
// @Produce  json
// @Success 200 {string} json "{"Code":0,"Msg":"","Body":"msg"}"
// @Router /Data/GetAccountInfo [get]
func GetAccountInfo(c *gin.Context) {
	c.JSON(http.StatusOK, models.ResultModel{
		Code: 0,
		Msg:  "",
		Body: "abc",
	})
}
