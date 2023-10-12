package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/waymon"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportAdd(c *gin.Context) {
	reportService := services.ReportService{}
	claims, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&reportService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := reportService.ReportAdd(c.Request.Context(), int64(claims.ID))
	c.JSON(http.StatusOK, res)
}
