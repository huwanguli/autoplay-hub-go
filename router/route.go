package router

import (
	"autoplay-hub/controller"
	"autoplay-hub/logger"
	"autoplay-hub/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由初始化
func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	{
		// 注册
		v1.POST("/register", controller.RegisterHandler)
		// 登录
		v1.POST("/login", controller.LoginHandler)
		// 使用JWT认证中间件
		v1.Use(middlewares.JWTAuthMiddleware())
		{
			// 创建脚本
			v1.POST("/script", controller.ScriptHandler)
			// 获取所有
			v1.GET("/scripts", controller.AllScriptInfoHandler)
			// 脚本详情
			v1.GET("/script/:id", controller.ScriptDetailHandler)
			// 编辑脚本
			v1.PATCH("/script/:id", controller.UpdateScriptHandler)
			// 删除脚本
			v1.DELETE("script/:id", controller.DeleteScriptHandler)
			// 运行脚本
			//v1.POST("/script/:id/run", controller.ScriptRunHandler)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
		})
	})
	return r
}
