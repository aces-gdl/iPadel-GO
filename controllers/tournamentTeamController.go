package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIsPlayerTaken(c *gin.Context) {

	TournamentIDstr := c.DefaultQuery("TournamentID", "0")
	CategoryIDstr := c.DefaultQuery("CategoryID", "0")
	//TeamIDstr := c.DefaultQuery("TeamID", "0")
	UserIDstr := c.DefaultQuery("UserID", "0")

	var team models.TournamentTeam
	whereStatement := `tournament_id=` + TournamentIDstr + ` and category_id = ` + CategoryIDstr + ` and ( member1_id = ` + UserIDstr + ` or member2_id = ` + UserIDstr + ` )`
	initializers.DB.Debug().Find(&team, whereStatement)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": team})

}

func PostEnrolledTeams(c *gin.Context) {
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
		Name1:        body.Name1,
		Ranking1:     body.Ranking1,
		Member2ID:    body.Member2ID,
		Name2:        body.Name2,
		Ranking2:     body.Ranking2,
		CategoryID:   body.CategoryID,
		TournamentID: body.TournamentID,
	}

	result := initializers.DB.Create(&tournamentTeam)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": tournamentTeam.ID})
}

func PutEnrolledTeams(c *gin.Context) {
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
		Name1:        body.Name1,
		Ranking1:     body.Ranking1,
		Member2ID:    body.Member2ID,
		Name2:        body.Name2,
		Ranking2:     body.Ranking2,
		CategoryID:   body.CategoryID,
		TournamentID: body.TournamentID,
	}
	tournamentTeam.ID = body.ID

	result := initializers.DB.Save(&tournamentTeam)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": tournamentTeam.ID})
}

func GetEnrolledTeams(c *gin.Context) {
	var tournamentIDstr = c.DefaultQuery("TournamentID", "0")
	var tournamentID, _ = strconv.ParseUint(tournamentIDstr, 10, 64)

	var categoryIDstr = c.DefaultQuery("CategoryID", "0")
	var categoryID, _ = strconv.ParseUint(categoryIDstr, 10, 64)

	var teams []models.TournamentTeam
	results := initializers.DB.Debug().Where("tournament_id = ? AND category_id = ?", tournamentID, categoryID).Find(&teams)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(teams), "data": teams})

}
