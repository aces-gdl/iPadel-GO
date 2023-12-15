package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func UploadImage(c *gin.Context) {
	img, err := imageupload.Process(c.Request, "file")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer imagen...",
		})
		return
	}
	id := c.Request.FormValue("ID")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer ID de imagen...",
		})
		return
	}

	fileType := strings.Split(img.ContentType, "/")[1]

	fileName := fmt.Sprintf("./pictures/%s.%s", id, fileType)
	thumbName := fmt.Sprintf("./pictures/%s-thumb.%s", id, fileType)

	regImage, err := imageupload.ThumbnailJPEG(img, 400, 400, 100)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al crear imagen...",
		})
		return
	}
	regImage.Save(fileName)

	thumb, err := imageupload.ThumbnailJPEG(img, 100, 100, 100)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al crear miniatura de imagen...",
		})
		return
	}
	thumb.Save(thumbName)

	c.JSON(http.StatusOK, gin.H{})
}

func GetImage(c *gin.Context) {
	id := c.Request.FormValue("ID")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer ID de imagen...",
		})
		return
	}

	c.Request.Header.Set("Content-Type", "image/png")
	if c.Request.FormValue("thumb") != "" {
		c.File(fmt.Sprintf("./pictures/%s-thumb.png", id))

	} else {
		c.File(fmt.Sprintf("./pictures/%s.png", id))

	}
	//c.JSON(http.StatusOK, gin.H{})
}
