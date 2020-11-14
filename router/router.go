package router

import (
	"github.com/gin-gonic/gin"
	"questionnaire_api/api/ctrl/controllers"
	"questionnaire_api/middleware/jwt"
)

func InitRouter(r *gin.Engine) {
	baseGroup := r.Group("api/v1/")

	//需要鉴权路由组
	baseAuthGroup := baseGroup.Group("", jwt.JWTAuthMiddleware())

	noAuth(baseGroup)

	auth(baseAuthGroup)

}

//不需要鉴权路由组
func noAuth(baseGroup *gin.RouterGroup) {
	authGroup := baseGroup.Group("auth")
	{
		authGroup.GET("", controllers.GetAccessToken)
		authGroup.POST("", controllers.AccessToken)
	}
}

//需要鉴权路由组
func auth(baseAuthGroup *gin.RouterGroup) {
	userGroup := baseAuthGroup.Group("user")
	{
		userGroup.GET("", controllers.GetUserInfo)
	}

	templateGroup := baseAuthGroup.Group("template")
	{
		templateGroup.POST("", controllers.AddTemplate())
		templateGroup.PUT("", controllers.UpdateTemplate())
		templateGroup.GET("",controllers.SearchTemplate())
	}
}
