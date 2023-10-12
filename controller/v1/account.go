package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/waymon"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountInfo(c *gin.Context) {
	accountService := services.AccountService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&accountService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := accountService.AccountInfo(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}

func AccountAdd(c *gin.Context) {
	accountService := services.AccountService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&accountService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := accountService.AccountAdd(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}
