package controllers

import (
	"fmt"
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostCreateGroups(c *gin.Context) {

	// Round Robin Formula
	// Games = Teams (Teams -1) / 2

	const groupSize int = 3

	var body struct {
		CategoryID   uint
		TournamentID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	var tournament models.Tournament

	initializers.DB.First(&tournament, "id= ?", body.TournamentID)
	if tournament.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Torneo no existe... ",
		})
		return
	}

	var teams []models.TournamentTeam
	// equipos de la categoria correcta
	// Cambiar query para que los ordene por el mas alto de cada pareja y arrastre al compañero
	sql := `select * 
				from tournament_teams tt 
				where tournament_id ='` + strconv.FormatUint(uint64(body.TournamentID), 10) + `' and category_id ='` + strconv.FormatUint(uint64(body.CategoryID), 10) + `' 
				order by  (case when ranking1 > ranking2 then ranking1 else ranking2 end) desc;`
	results := initializers.DB.Debug().Raw(sql).Find(&teams)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuarios no existen... ",
		})
		return
	}

	teamCounter := int(results.RowsAffected)

	groupsCounter := teamCounter / groupSize

	groupsCounterFinal := int(groupsCounter)
	if (teamCounter % groupSize) != 0 {
		groupsCounterFinal = int(groupsCounter) + 1
	}
	groups := make([]struct {
		ID      uint
		counter int
	}, groupsCounterFinal)

	for i := 0; i < groupsCounterFinal; i++ {
		var group models.TournamentGroup
		fmt.Println("group ", i+1)
		group.Name = fmt.Sprintf("Grupo - %02d", i+1)
		group.TournamentID = body.TournamentID
		group.CategoryID = body.CategoryID
		group.GroupNumber = i + 1
		result := initializers.DB.Create(&group)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Fallo al crear grupo... ",
			})
			return
		}
		groups[i].ID = group.ID
		groups[i].counter = 1

	}
	goingUp := true
	teamSelector := 0
	groupSelector := 0
	teamCounterByGroup := 1
	for {

		fmt.Println("Group : ", groupSelector, "  --> Team : ", teamSelector+1)
		var teamByGroup models.TournamentTeamByGroup
		teamByGroup.TournamentID = body.TournamentID
		teamByGroup.CategoryID = body.CategoryID
		teamByGroup.GroupID = groups[groupSelector].ID
		teamByGroup.GroupNumber = groupSelector + 1
		teamByGroup.Name = fmt.Sprintf("Equipo : %02d", teamCounterByGroup)
		teamByGroup.TeamID = teams[teamSelector].ID
		initializers.DB.Create(&teamByGroup)
		if goingUp {
			groupSelector++
		} else {
			groupSelector--
		}
		teamSelector++
		if groupSelector < 0 || groupSelector > groupsCounter-1 {
			goingUp = !goingUp
			if groupSelector < 0 {
				groupSelector = 0
			}
			if groupSelector > groupsCounter-1 {
				groupSelector = groupsCounter - 1
			}
			teamCounterByGroup++
			if teamCounterByGroup > groupSize {
				teamCounterByGroup = 1
			}
		}
		if teamSelector >= int(groupsCounter*groupSize) {
			break
		}
	}
	if teamSelector < int(teamCounter) {
		teamCounterByGroup = 1
		for i := teamSelector; i < int(teamCounter); i++ {
			fmt.Println(" Extra Group : ", groupsCounterFinal, "  --> Team : ", i+1)
			var teamByGroup models.TournamentTeamByGroup
			teamByGroup.TournamentID = body.TournamentID
			teamByGroup.CategoryID = body.CategoryID
			teamByGroup.GroupID = groups[groupsCounterFinal-1].ID
			teamByGroup.GroupNumber = groupsCounterFinal
			teamByGroup.Name = fmt.Sprintf("Equipo : %02d", teamCounterByGroup)
			teamByGroup.TeamID = teams[i].ID
			initializers.DB.Create(&teamByGroup)
			teamCounterByGroup++
		}
	}

	// Crear Juegos

	type Game struct {
		Team1 int
		Team2 int
	}

	var roleOfGames []Game

	AddNewGame := func(team1 int, team2 int) {
		var gameNotFound = true
		for i := 0; i < len(roleOfGames); i++ {
			if (roleOfGames[i].Team1 == team1 && roleOfGames[i].Team2 == team2) || (roleOfGames[i].Team1 == team2 && roleOfGames[i].Team2 == team1) {
				gameNotFound = false
				break
			}
		}
		if gameNotFound {
			var newGame = Game{Team1: team1, Team2: team2}
			roleOfGames = append(roleOfGames, newGame)
		}
	}

	CreateGames := func(TournamentID uint, CategoryID uint, teams []models.TournamentTeamByGroup) {
		var numOfTeams int = len(teams)
		roleOfGames = nil

		for i := 0; i < numOfTeams; i++ {
			for x := 0; x < numOfTeams; x++ {
				if i != x {
					AddNewGame(i, x)
				}
			}
		}

		for i := 0; i < len(roleOfGames); i++ {
			fmt.Println("Game ", i+1, " --> Team 1: ", roleOfGames[i].Team1, " vs Team 2:", roleOfGames[i].Team2)
			var newGame models.TournamentGames
			newGame.TournamentID = TournamentID
			newGame.CategoryID = CategoryID
			newGame.TournamentGroupID = teams[roleOfGames[i].Team1].GroupID
			newGame.Team1ID = teams[roleOfGames[i].Team1].TeamID
			newGame.Team2ID = teams[roleOfGames[i].Team2].TeamID
			newGame.GameType = "RoundRobin"

			initializers.DB.Create(&newGame)
		}
	}

	var tournamentGroup []models.TournamentGroup
	results = initializers.DB.Where("tournament_id = ? AND category_id = ?", body.TournamentID, body.CategoryID).Order("group_number").Find(&tournamentGroup)
	if results.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Grupos no existen... ",
		})
		return
	}

	var teamsByGroup []models.TournamentTeamByGroup
	for i := 0; i < len(tournamentGroup); i++ {
		fmt.Println("Partidos grupo : ", tournamentGroup[i].GroupNumber)
		results = initializers.DB.Where("group_id = ?", tournamentGroup[i].ID).Find(&teamsByGroup)
		CreateGames(body.TournamentID, body.CategoryID, teamsByGroup)
	}

}
