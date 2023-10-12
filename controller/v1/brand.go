package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BrandList(c *gin.Context) {
	brandService := services.BrandService{}
	if err := c.ShouldBind(&brandService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := brandService.BrandList(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
