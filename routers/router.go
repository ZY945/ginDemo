package routers

import (
	"GinAndSqlx/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自定义中间件
func MyHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		//拦截器逻辑
		context.Next() //放行
		//context.Abort() //阻止
	}
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(MyHandle())
	userEngine := router.Group("/user")
	{
		userEngine.GET("", controllers.GetUserVoById)
		userEngine.GET("list", controllers.List)
		userEngine.POST("", controllers.InsertUser)
		userEngine.PUT("", controllers.UpdateUser)
		userEngine.DELETE(":id", controllers.DelUser)
		userEngine.GET("login", controllers.Login)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return router
}
