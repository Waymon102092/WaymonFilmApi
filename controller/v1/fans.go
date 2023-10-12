package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/waymon"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FansList(c *gin.Context) {
	fansService := services.FansService{}
	claims, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&fansService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := fansService.FansList(c.Request.Context(), int64(claims.ID))
	c.JSON(http.StatusOK, res)
}

func FansCount(c *gin.Context) {
	fansService := services.FansService{}
	claims, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&fansService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := fansService.FansCount(c.Request.Context(), int64(claims.ID))
	c.JSON(http.StatusOK, res)
}

func FansOrderCount(c *gin.Context) {
	fansService := services.FansService{}
	claims, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&fansService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := fansService.FansOrderCount(c.Request.Context(), int64(claims.ID))
	c.JSON(http.StatusOK, res)
}

func FansOrder(c *gin.Context) {
	fansService := services.FansService{}
	claims, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&fansService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := fansService.FansOrder(c.Request.Context(), int64(claims.ID))
	c.JSON(http.StatusOK, res)
}
