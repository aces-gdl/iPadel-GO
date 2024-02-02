// Package Classification of Games API
//
// Documentation for Games API
//
//  Schemes: http
//  BasePath: /v1/tournament/listgames
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta

package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetGames(c *gin.Context) {

	var TournamentID = c.DefaultQuery("TournamentID", "")
	var SearchStr = c.DefaultQuery("SearchStr", "")

	type gameList struct {
		TournamentID          uuid.UUID
		CategoryID            uuid.UUID
		TournamentGroupID     uuid.UUID
		Team1ID               uuid.UUID
		Team2ID               uuid.UUID
		GameID                uuid.UUID
		Team1Member1ID        uuid.UUID
		Team1Name1            string
		Team1Ranking1         int
		Team1Member2ID        uuid.UUID
		Team1Name2            string
		Team1Ranking2         int
		Team2Member1ID        uuid.UUID
		Team2Name1            string
		Team2Ranking1         int
		Team2Member2ID        uuid.UUID
		Team2Name2            string
		Team2Ranking2         int
		GroupNumber           int
		TournamentTimeSlotsID uuid.UUID
		CategoryColor         string
		CategoryDescription   string
		GameResultsID         uuid.UUID
		Team1Set1             int
		Team1Set2             int
		Team1Set3             int
		Team2Set1             int
		Team2Set2             int
		Team2Set3             int
	}

	var gameLists []gameList
	unassignedGameQuery := `select
								tg.ID as game_id,
								tg.tournament_id ,
								tgrp.group_number ,
								tgrp.id as tournament_group_id,
								c.id as category_id,
								c.color as category_color,
								c.description as category_description,  
								tt1.member1_id as team1_member1_id,
								tt1.name1 as team1_name1,
								tt1.ranking1 as team1_ranking1,
								tt1.member2_id as team1_member2_id,
								tt1.name2 as team1_name2,
								tt1.ranking2 as team1_ranking2,
								tt2.member1_id as team2_member1_id,
								tt2.name1 as team2_name1,
								tt2.ranking1 as team2_ranking1,
								tt2.member2_id as team2_member2_id,
								tt2.name2 as team2_name2,
								tt2.ranking2 as team2_ranking2,
								tg.tournament_time_slots_id,
								gr.id as game_results_id,
								gr.team1_set1,
								gr.team1_set2,
								gr.team1_set3, 
								gr.team2_set1,
								gr.team2_set2,
								gr.team2_set3 
							from tournament_games tg 
								inner join tournament_teams tt1 on tg.team1_id = tt1.id and tg.tournament_id = tt1.tournament_id 
								inner join tournament_teams tt2 on tg.team2_id = tt2.id and tg.tournament_id = tt2.tournament_id 
								inner join categories c on tg.category_id = c.id 
								inner join tournament_groups tgrp on tg.tournament_group_id = tgrp.id
							    left outer join tournament_game_results gr on tg.id = gr.game_id
							where tg.tournament_id ='` + TournamentID + `' `
	if SearchStr != "" {
		unassignedGameQuery += ` and ( tt1.name1 like '%` + SearchStr + `%' or tt1.name2 like '%` + SearchStr + `%' or tt2.name1 like '%` + SearchStr + `%' or tt2.name2 like '%` + SearchStr + `%' ) `
	}
	unassignedGameQuery += `	order by tgrp.group_number`
	results := initializers.DB.Debug().Raw(unassignedGameQuery).Scan(&gameLists)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(gameLists), "data": gameLists})

}

func PutAssignGamesToTimeSlots(c *gin.Context) {
	var body struct {
		TournamentID string
		GameID       string
		TimeSlotID   string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	TournamentID, err := uuid.Parse(body.TournamentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error})
		return
	}

	// Validate TournamentID
	var tournament models.Tournament
	results := initializers.DB.First(&tournament, TournamentID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": results.Error})
		return
	}

	GameID, err := uuid.Parse(body.GameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error})
		return
	}
	// Validate GameID
	var game models.TournamentGames
	results = initializers.DB.First(&game, GameID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": results.Error})
		return
	}

	TimeSlotID, err := uuid.Parse(body.TimeSlotID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error})
		return
	}
	// Validate TimeSlotsID
	var timeSlot models.TournamentTimeSlots
	results = initializers.DB.First(&timeSlot, TimeSlotID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": results.Error})
		return
	}

	game.TournamentTimeSlotsID = timeSlot.ID
	results = initializers.DB.Debug().Save(&game)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	timeSlot.GameID = game.ID
	results = initializers.DB.Debug().Save(&timeSlot)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": 1, "data": game})

}

func DeleteAssignGamesToTimeSlots(c *gin.Context) {
	var body struct {
		TournamentID string
		GameID       string
		TimeSlotID   string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	TournamentID, err := uuid.Parse(body.TournamentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error})
		return
	}

	// Validate TournamentID
	var tournament models.Tournament
	results := initializers.DB.First(&tournament, TournamentID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": results.Error})
		return
	}

	GameID, err := uuid.Parse(body.GameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error})
		return
	}
	// Validate GameID
	var game models.TournamentGames
	results = initializers.DB.First(&game, GameID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": results.Error})
		return
	}

	TimeSlotID, err := uuid.Parse(body.TimeSlotID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error})
		return
	}
	// Validate TimeSlotsID
	var timeSlot models.TournamentTimeSlots
	results = initializers.DB.First(&timeSlot, TimeSlotID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": results.Error})
		return
	}

	game.TournamentTimeSlotsID = uuid.Nil
	results = initializers.DB.Save(&game)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	timeSlot.GameID = uuid.Nil
	results = initializers.DB.Save(&timeSlot)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": 1, "data": game})

}

func PostCreateGame(c *gin.Context) {

	var body struct {
		TournamentID string
		CategoryID   string
		GroupID      string
		Team1ID      string
		Team2ID      string
		Comment      string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error})
		return
	}
	var newGame models.TournamentGames
	TournamentID, _ := uuid.Parse(body.TournamentID)
	CategoryID, _ := uuid.Parse(body.CategoryID)
	GroupID, _ := uuid.Parse(body.GroupID)
	Team1ID, _ := uuid.Parse(body.Team1ID)
	Team2ID, _ := uuid.Parse(body.Team2ID)

	// Validate TournamentID
	var tournament models.Tournament
	results := initializers.DB.Debug().First(&tournament, TournamentID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Tournament not found..."})
		return
	}

	//Validate Category
	var category models.Category
	results = initializers.DB.Debug().First(&category, CategoryID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Category not found..."})
		return
	}

	var Team1 models.TournamentTeam
	results = initializers.DB.Debug().First(&Team1, Team1ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Team1 not found ..."})
		return
	}

	var Team2 models.TournamentTeam
	results = initializers.DB.Debug().First(&Team2, Team2ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Team2 not found ..."})
		return
	}

	newGame.TournamentID = TournamentID
	newGame.CategoryID = CategoryID
	newGame.TournamentGroupID = GroupID
	newGame.Team1ID = Team1ID
	newGame.Team2ID = Team2ID
	newGame.Comment = body.Comment
	newGame.Active = true

	results = initializers.DB.Debug().Create(&newGame)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": 1, "data": newGame})
}
