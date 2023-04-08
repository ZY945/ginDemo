package routers

import "github.com/gin-gonic/gin"
import "GinAndSqlx/controllers"

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
	}

	return router
}
