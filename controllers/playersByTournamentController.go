package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostRegisterUserForTournament(c *gin.Context) {
	var body models.PlayersByTournament

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	playerForTournament := models.PlayersByTournament{
		TournamentID:  body.TournamentID,
		CategoryID:    body.CategoryID,
		UserID:        body.UserID,
		PaymentStatus: "Pendiente",
	}
	initializers.DB.Debug().Create(&playerForTournament)

	c.JSON(http.StatusOK, gin.H{"ID": playerForTournament.ID})
}

func GetRegisteredUsersByTournament(c *gin.Context) {
	var categoryID = c.DefaultQuery("CategoryID", "")
	var tournamentID = c.DefaultQuery("TournamentID", "")
	var SearchString = c.DefaultQuery("SearchString", "")

	var whereClause = " where 1 = 1 "
	if categoryID != "" {
		whereClause = whereClause + " AND u.category_id = " + categoryID
	}
	if tournamentID != "" {
		whereClause = whereClause + " AND pbt.tournament_id = " + tournamentID
	}
	if SearchString != "" {
		whereClause = whereClause + " AND u.name like  '%" + SearchString + "%'"
	}

	queryString := `SELECT u.*, 
		c.description as category_description, 
		c.color as category_color,
		pbt.payment_status 
	FROM "users" u
		left join categories c on u.category_id = c.id 
		left join players_by_tournaments pbt on u.id = pbt.user_id  and c.id = pbt.category_id 		
 ` + whereClause + `
		Order by  u.name asc`

	type userExtended struct {
		ID                    uint
		CategoryID            uint
		CategoryColor         string
		Name                  string
		FamilyName            string
		GivenName             string
		Email                 string
		CategoryDescription   string
		PermissionID          uint
		Phone                 string
		HasPicture            int
		MemberSince           time.Time
		Birthday              time.Time
		PermissionDescription string
		Ranking               int
		PaymentStatus         string
	}

	var usersExtended []userExtended

	results := initializers.DB.Debug().Raw(queryString).Where(whereClause).Scan(&usersExtended)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(usersExtended), "data": usersExtended})
}
