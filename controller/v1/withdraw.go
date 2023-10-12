package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/waymon"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WithdrawInfo(c *gin.Context) {
	withdrawService := services.WithdrawService{}
	if err := c.ShouldBind(&withdrawService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := withdrawService.WithdrawInfo(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func WithdrawList(c *gin.Context) {
	withdrawService := services.WithdrawService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&withdrawService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := withdrawService.WithdrawList(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}

func WithdrawAdd(c *gin.Context) {
	withdrawService := services.WithdrawService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&withdrawService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := withdrawService.WithdrawAdd(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}

func WithdrawEdit(c *gin.Context) {
	withdrawService := services.WithdrawService{}
	if err := c.ShouldBind(&withdrawService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := withdrawService.WithdrawEdit(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func WithdrawStatus(c *gin.Context) {
	withdrawService := services.WithdrawService{}
	if err := c.ShouldBind(&withdrawService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := withdrawService.WithdrawStatus(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func WithdrawMoney(c *gin.Context) {
	withdrawService := services.WithdrawService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&withdrawService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := withdrawService.WithdrawMoney(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}
