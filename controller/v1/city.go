package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CityAdd(c *gin.Context) {
	cityService := services.CityService{}
	if err := c.ShouldBind(&cityService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := cityService.CityAdd(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func CityCurrent(c *gin.Context) {
	cityService := services.CityService{}
	if err := c.ShouldBind(&cityService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := cityService.CityCurrent(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
