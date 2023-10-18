package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostCourts(c *gin.Context) {
	var body models.Court

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	courts := models.Court{
		ID:      body.ID,
		Name:    body.Name,
		Indoors: body.Indoors,
		ClubID:  body.ClubID,
	}

	result := initializers.DB.Create(&courts)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetCourts(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var courts []models.Court
	results := initializers.DB.Limit(intLimit).Offset(offset).Find(&courts)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(courts), "data": courts})
}

func GetCourtsByClub(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")
	var ClubID = c.DefaultQuery("ClubID", "")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var courts []models.Court
	results := initializers.DB.Limit(intLimit).Offset(offset).Where("club_id = ?", ClubID).Find(&courts)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(courts), "data": courts})
}
