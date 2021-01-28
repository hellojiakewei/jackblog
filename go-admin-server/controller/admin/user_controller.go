package admin

import (
	"github.com/gin-gonic/gin"
	"jackblog/model/response"
)

func Login(c *gin.Context) {

		response.OkWithMessage("nihao", c)
	}
