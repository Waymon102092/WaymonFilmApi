package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DistrictFilter(c *gin.Context) {
	districtService := services.DistrictService{}
	if err := c.ShouldBind(&districtService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := districtService.DistrictFilter(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
