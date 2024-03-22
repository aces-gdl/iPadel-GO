package controllers

import (
	"iPadel-GO/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	var tournamentID = c.DefaultQuery("TournamentID", "")
	var categoryID = c.DefaultQuery("CategoryID", "")

	type groupExtended struct {
		TournamentID uint
		CategoryID   uint
		GroupID      uint
		GroupNumber  int
		Name         string
		Name1        string
		Ranking1     int
		Name2        string
		Ranking2     int
		TeamRanking  int
	}
	var groupsExtended []groupExtended

	groupsQuery := `select 
					ttbg.tournament_id , 
					ttbg.category_id , 
					ttbg.group_id, 
					ttbg.group_number , 
					tt.name, 
					tt.name1, 
					tt.ranking1 , 
					tt.name2,
					tt.ranking2,
					tt.ranking1 + tt.ranking2 as team_ranking
					from tournament_team_by_groups ttbg inner join tournament_teams tt on ttbg.team_id  = tt.id
					where ttbg.category_id = '` + categoryID + `' and ttbg.tournament_id = '` + tournamentID + `'
					order by ttbg.group_number ASC, tt.ranking1 + ranking2 DESC`

	results := initializers.DB.Raw(groupsQuery).Where("tournament_id = ? AND category_id = ?", tournamentID, categoryID).Scan(&groupsExtended)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(groupsExtended), "data": groupsExtended})

}
