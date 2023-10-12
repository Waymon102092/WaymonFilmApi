package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MovieHot(c *gin.Context) {
	movieService := services.MovieService{}
	if err := c.ShouldBind(&movieService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := movieService.MovieHot(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func MovieComing(c *gin.Context) {
	movieService := services.MovieService{}
	if err := c.ShouldBind(&movieService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := movieService.MovieComing(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
