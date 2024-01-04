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
								tg.tournament_time_slots_id 
							from tournament_games tg 
								inner join tournament_teams tt1 on tg.team1_id = tt1.id and tg.tournament_id = tt1.tournament_id 
								inner join tournament_teams tt2 on tg.team2_id = tt2.id and tg.tournament_id = tt2.tournament_id 
								inner join categories c on tg.category_id = c.id 
								inner join tournament_groups tgrp on tg.tournament_group_id = tgrp.id
							where tg.tournament_id ='` + TournamentID + `'  `
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
	results = initializers.DB.Save(&game)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	timeSlot.GameID = game.ID
	results = initializers.DB.Save(&timeSlot)
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
