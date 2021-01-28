package initconf

import (
	"github.com/gin-gonic/gin"
	"jackblog/conf"
	"jackblog/mid"
	"jackblog/router"
)

func InitRouter()  {
	var r =gin.Default()
	r.Use(mid.CorsMid())
	admin:=r.Group("admin")
	{
		router.InitUserRouter(admin)
	}
	r.Run(conf.WebPort)
}
