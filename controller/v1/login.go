package v1

import (
	"Waymon_api/pkg/e"
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Author(c *gin.Context) {
	authorService := services.AuthorService{}
	if err := c.ShouldBind(&authorService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := authorService.Author(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func Session(c *gin.Context) {
	authorService := services.AuthorService{}
	if err := c.ShouldBind(&authorService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := authorService.Session(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func Login(c *gin.Context) {
	loginService := services.LoginService{}
	if err := c.ShouldBind(&loginService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := loginService.Login(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func TTLogin(c *gin.Context) {
	loginService := services.LoginService{}
	if err := c.ShouldBind(&loginService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := loginService.TTLogin(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {
	loginService := services.LoginService{}
	if err := c.ShouldBind(&loginService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := loginService.Register(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func TTRegister(c *gin.Context) {
	loginService := services.LoginService{}
	if err := c.ShouldBind(&loginService); err != nil {
		code := e.ParamError
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   nil,
		})
		return
	}
	res := loginService.TTRegister(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
