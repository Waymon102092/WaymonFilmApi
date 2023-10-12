package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CustomInfo(c *gin.Context) {
	customService := services.CustomService{}
	if err := c.ShouldBind(&customService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := customService.CustomInfo(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
