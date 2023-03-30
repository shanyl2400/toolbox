package api

import (
	"net/http"
	"strings"
	"toolbox/internal/utils"

	"github.com/gin-gonic/gin"
)

func (api *API) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//cross domain
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT,DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (api *API) checkAuth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusForbidden, GeneralResponse{
			Message: "token为空",
		})
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusForbidden, GeneralResponse{
			Message: "请求头中auth格式有误",
		})
		c.Abort()
		return
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	mc, err := utils.ParseToken(parts[1])
	if err != nil {
		c.JSON(http.StatusForbidden, GeneralResponse{
			Message: "无效的Token",
		})
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	c.Set("username", mc.Username)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}
