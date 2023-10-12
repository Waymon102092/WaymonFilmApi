package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelpInfo(c *gin.Context) {
	helpService := services.HelpService{}
	if err := c.ShouldBind(&helpService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := helpService.HelpInfo(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func HelpList(c *gin.Context) {
	helpService := services.HelpService{}
	if err := c.ShouldBind(&helpService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := helpService.HelpList(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
