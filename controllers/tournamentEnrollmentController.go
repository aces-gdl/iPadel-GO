package controllers

import (
	"fmt"
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostSimulateEnrollment(c *gin.Context) {
	var body struct {
		CategoryID   uint
		UserCount    string
		TournamentID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	// Buscar Usuarios
	userCount, _ := strconv.Atoi(body.UserCount)
	var users []models.User
	results := initializers.DB.Debug().Where("category_id = ?", body.CategoryID).Preload(clause.Associations).Limit(userCount).Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	var foundUsers int = int(results.RowsAffected)
	fmt.Println(foundUsers)

	// insertar en teams los 2 usuarios
	counter := len(users) / 2

	for i := 0; i < counter; i++ {
		var team models.TournamentTeam
		var teamMember1 models.User
		var teamMember2 models.User

		team.Member1ID = users[i+1].ID
		initializers.DB.Debug().First(&teamMember1, "id= ?", team.Member1ID)
		if teamMember1.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuario 1 no encontrado... ",
			})
			return
		}
		team.Member2ID = users[i+counter].ID
		initializers.DB.Debug().First(&teamMember2, "id= ?", team.Member2ID)
		if teamMember2.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuario 2 no encontrado... ",
			})
			return
		}

		team.Name = fmt.Sprintf("Pareja  - %02d", i+1)
		team.TournamentID = body.TournamentID
		team.CategoryID = body.CategoryID

		team.Name1 = teamMember1.Name
		team.Ranking1 = teamMember1.Ranking

		team.Name2 = teamMember2.Name
		team.Ranking2 = teamMember2.Ranking

		fmt.Println(team)
		result := initializers.DB.Debug().Create(&team)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Fallo al crear Equipo... ",
			})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{})
}
