package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostClub(c *gin.Context) {
	var body models.Club

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	club := models.Club{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Contact:     body.Contact,
		Phone:       body.Phone,
		Address:     body.Address,
	}

	result := initializers.DB.Create(&club)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetClubs(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var clubs []models.Club
	results := initializers.DB.Preload(clause.Associations).Limit(intLimit).Offset(offset).Find(&clubs)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(clubs), "data": clubs})
}
