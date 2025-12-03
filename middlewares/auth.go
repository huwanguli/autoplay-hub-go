package middlewares

import (
	"autoplay-hub/controller"
	"autoplay-hub/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// JWTAuthMiddleware Gin 中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			zap.L().Error("Token为空")
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			zap.L().Error("Token解析错误", zap.Error(err))
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set(controller.CtxUserIDKey, claims.UserID)
		c.Set(controller.CtxIsAdminKey, claims.IsAdmin)
		c.Next()
	}
}
