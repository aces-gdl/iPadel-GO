package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostCategory(c *gin.Context) {

	var body struct {
		Description string
		Level       int
		Active      bool
		Color       string
	}

	x := c.Bind(&body)
	if x != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	//level, _ := strconv.Atoi(body.Level)
	//active, _ := strconv.ParseBool(body.Active)
	category := models.Category{
		Description: body.Description,
		Level:       body.Level,
		Active:      body.Active,
		Color:       body.Color,
	}

	result := initializers.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetCatgories(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	CategoryIDStr := c.DefaultQuery("ID", "")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var categories []models.Category

	SQLStr := "Select * from categories where 1 = 1 "
	if CategoryIDStr != "" {
		SQLStr = SQLStr + ` and id = ` + CategoryIDStr
	}

	SQLStr += "order by level asc, description asc "
	results := initializers.DB.Raw(SQLStr).Limit(intLimit).Offset(offset).Scan(&categories)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(categories), "data": categories})
}
