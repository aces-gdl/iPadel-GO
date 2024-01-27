package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRoundRobinWinner(c *gin.Context) {
	tournamentID := c.DefaultQuery("TournamentID", "")

	var tournament models.Tournament

	result := initializers.DB.First(&tournament, "id = ?", tournamentID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Torneo no encontrado... ",
		})
		return
	}
	categoryID := c.DefaultQuery("CategoryID", "")

	var category models.Category

	result = initializers.DB.First(&category, "id = ?", categoryID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Categoria no encontrada... ",
		})
		return
	}

	groupID := c.DefaultQuery("GroupID", "")
	if groupID != "" {
		var group models.TournamentGroup

		result = initializers.DB.First(&group, "id = ?", groupID)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Grupo no encontrada... ",
			})
			return
		}
	}

	type gameResultsExtended struct {
		GameID         uuid.UUID
		TournamentID   uuid.UUID
		CategoryID     uuid.UUID
		GroupID        uuid.UUID
		GroupName      string
		Team1ID        uuid.UUID
		Team2ID        uuid.UUID
		Team1Member1   string
		Team1Member1ID uuid.UUID
		Team1Member2   string
		Team1Member2ID uuid.UUID
		Team2Member1   string
		Team1Member3ID uuid.UUID
		Team2Member2   string
		Team1Member4ID uuid.UUID
		Team1Set1      int
		Team1Set2      int
		Team1Set3      int
		Team2Set1      int
		Team2Set2      int
		Team2Set3      int
	}

	var gamesResultsExtended []gameResultsExtended
	queryStr := `select 
	tg.id as game_id,
	tg.tournament_id,
	tg.category_id ,
	tg2."name" as group_name,
	tg.tournament_group_id as group_id,
	tg.team1_id,
	tg.team2_id ,
	us1.name as team1_member1,
	us1.id as team1_member1_id,
	us2.name as team1_member2,
	us2.id as team1_member2_id,
	us3.name as team2_member1,
	us3.id as team1_member3_id,
	us4.name as team2_member2,
	us4.id as team1_member4_id,
	tgr.team1_set1,
	tgr.team1_set2 ,
	tgr.team1_set3 ,
	tgr.team2_set1 ,
	tgr.team2_set2 ,
	tgr.team2_set3 
from tournament_games tg
    left outer join tournament_game_results tgr on tg.id = tgr.game_id
    inner join tournament_teams tt1 on tg.team1_id = tt1.id 
    inner join tournament_teams tt2 on tg.team2_id = tt2.id 
    inner join users us1 on tt1.member1_id = us1.id 
    inner join users us2 on tt1.member2_id = us2.id 
    inner join users us3 on tt2.member1_id = us3.id 
    inner join users us4 on tt2.member2_id = us4.id 
    inner join tournament_groups tg2 on tg.tournament_group_id = tg2.id
where
	tg.tournament_id = '` + tournamentID + `' AND tg.category_id = '` + categoryID + `'
order by tg2."name"`

	if groupID != "" {
		queryStr += ` AND tg.tournament_group_id = '` + groupID + `'`
	}

	result = initializers.DB.Debug().Raw(queryStr).Scan(&gamesResultsExtended)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Resultados de parttidos no encontrados... ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(gamesResultsExtended), "data": gamesResultsExtended})
}
