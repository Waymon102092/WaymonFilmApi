package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/waymon"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AmountList(c *gin.Context) {
	amountService := services.AmountService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&amountService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := amountService.AmountList(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}

func AmountMoney(c *gin.Context) {
	amountService := services.AmountService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&amountService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := amountService.AmountMoney(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}

func AmountSettle(c *gin.Context) {
	amountService := services.AmountService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&amountService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := amountService.AmountSettle(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}

func AmountAccumulate(c *gin.Context) {
	amountService := services.AmountService{}
	claim, _ := waymon.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&amountService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := amountService.AmountAccumulate(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, res)
}
