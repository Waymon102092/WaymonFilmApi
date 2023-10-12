package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CinemaList(c *gin.Context) {
	cinemaService := services.CinemaService{}
	if err := c.ShouldBind(&cinemaService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := cinemaService.CinemaList(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
