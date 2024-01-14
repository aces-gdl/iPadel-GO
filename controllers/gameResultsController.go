package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostGameResults(c *gin.Context) {
	var body struct {
		GameID        string
		Team1Set1     int
		Team1Set2     int
		Team1Set3     int
		Team2Set1     int
		Team2Set2     int
		Team2Set3     int
		Winner        int
		WinningReason string
		Comments      string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	GameID, _ := uuid.Parse(body.GameID)

	var gameResult = models.TournamentGameResult{
		GameID:        GameID,
		Team1Set1:     body.Team1Set1,
		Team1Set2:     body.Team1Set2,
		Team1Set3:     body.Team1Set3,
		Team2Set1:     body.Team2Set1,
		Team2Set2:     body.Team2Set2,
		Team2Set3:     body.Team2Set3,
		Winner:        body.Winner,
		WinningReason: body.WinningReason,
		Comments:      body.Comments,
	}

	results := initializers.DB.Debug().Create(&gameResult)
	if results.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": gameResult.ID,
	})
}
