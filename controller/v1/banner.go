package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BannerList(c *gin.Context) {
	bannerService := services.BannerService{}
	if err := c.ShouldBind(&bannerService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := bannerService.BannerList(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
