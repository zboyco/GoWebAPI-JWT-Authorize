package routers

import (
	"GoWebAPI-JWT-Authorize/controllers"
	"GoWebAPI-JWT-Authorize/middleware/jwt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 初始化Router
func InitRouter() *gin.Engine {
	r := gin.Default()

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 登录
	r.GET("/Login", controllers.Login)

	m := r.Group("/Data")
	m.Use(jwt.Auth())
	{
		m.GET("/GetAccountInfo", controllers.GetAccountInfo)
	}

	return r
}
