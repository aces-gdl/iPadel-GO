package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPermissions(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var permissions []models.Permission
	results := initializers.DB.Order("description asc").Limit(intLimit).Offset(offset).Find(&permissions)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(permissions), "data": permissions})
}

func PostPermissions(c *gin.Context) {
	var body struct {
		Description string
		Active      bool
	}

	x := c.Bind(&body)
	if x != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	//active, _ := strconv.ParseBool(body.Active)
	permission := models.Permission{
		Description: body.Description,
		Active:      body.Active,
	}

	result := initializers.DB.Create(&permission)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear permiso... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
