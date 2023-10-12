package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/waymon"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PosterList(c *gin.Context) {
	posterService := services.PosterService{}
	if err := c.ShouldBind(&posterService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := posterService.PosterList(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func PosterCode(c *gin.Context) {
	posterService := services.PosterService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&posterService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := posterService.PosterCode(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}
