package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigInfo(c *gin.Context) {
	configService := services.ConfigService{}
	if err := c.ShouldBind(&configService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := configService.ConfigInfo(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func ConfigEdit(c *gin.Context) {
	configService := services.ConfigService{}
	if err := c.ShouldBind(&configService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := configService.ConfigEdit(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
