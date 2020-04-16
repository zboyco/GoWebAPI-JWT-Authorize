package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zboyco/GoWebAPI-JWT-Authorize/utils"
	"log"
	"net/http"
)

// JwtAuth 中间件，检查token
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// parseToken 解析token包含的信息
		claims, err := utils.ParseToken(token)
		if err != nil {
			if err == utils.ErrorExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": -1,
					"msg":  "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("UserID", claims.UserID)
		// c.Next()
	}
}
