package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportCategoryList(c *gin.Context) {
	reportCategoryService := services.ReportCategoryService{}
	if err := c.ShouldBind(&reportCategoryService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := reportCategoryService.ReportCategoryList(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func ReportCategoryAdd(c *gin.Context) {
	reportCategoryService := services.ReportCategoryService{}
	if err := c.ShouldBind(&reportCategoryService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := reportCategoryService.ReportCategoryAdd(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
