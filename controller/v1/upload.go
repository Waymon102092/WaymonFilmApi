package v1

import (
	"Waymon_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	uploadService := services.UploadService{}
	res := uploadService.Upload(c.Request.Context(), files)
	c.JSON(http.StatusOK, res)
}
