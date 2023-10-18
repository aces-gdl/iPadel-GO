package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTeam(c *gin.Context) {
	var body models.TournamentTeam

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	tournamentTeam := models.TournamentTeam{
		Name:         body.Name,
		Member1ID:    body.Member1ID,
		Member2ID:    body.Member2ID,
		TournamentID: body.TournamentID,
	}

	result := initializers.DB.Create(&tournamentTeam)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
