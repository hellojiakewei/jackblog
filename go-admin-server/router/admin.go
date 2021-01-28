package router

import (
	"github.com/gin-gonic/gin"
	"jackblog/controller/admin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", admin.Login)       //登录
	}

}