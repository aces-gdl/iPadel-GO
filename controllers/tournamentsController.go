package controllers

import (
	"fmt"
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func GetTournaments(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var tournaments []models.Tournament
	results := initializers.DB.Preload(clause.Associations).Limit(intLimit).Offset(offset).Find(&tournaments)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(tournaments), "data": tournaments})
}

func ParseTime(timeOnlySrt string) time.Time {
	timeArray := strings.Split(timeOnlySrt, ":")
	hour, _ := strconv.Atoi(timeArray[0])
	minute, _ := strconv.Atoi(timeArray[1])
	myLocation, _ := time.LoadLocation("America/Mexico_City")
	result := time.Date(2023, 10, 02, hour, minute, 0, 0, myLocation)

	return result
}

func PostTournaments(c *gin.Context) {
	//var body models.Tournament

	var body struct {
		Description      string
		StartDate        string
		EndDate          string
		StartTime        string
		EndTime          string
		HostClubID       string
		GameDuration     string
		RoundrobinDays   string
		PlayOffDays      string
		RoundrobinCourts string
		PlayoffCourts    string
		Active           bool
	}
	resultTest := c.Bind(&body)
	if resultTest != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", body.StartDate)
	if err != nil {
		fmt.Println(err)
	}
	endDate, _ := time.Parse("2006-01-02", body.EndDate)
	clubID, _ := uuid.Parse(body.HostClubID)
	gameDuration, _ := strconv.Atoi(body.GameDuration)
	roundRobinCourts, _ := strconv.Atoi(body.RoundrobinCourts)
	playoffCourts, _ := strconv.Atoi(body.PlayoffCourts)
	roundrobinDays, _ := strconv.Atoi(body.RoundrobinDays)
	playoffDays, _ := strconv.Atoi(body.PlayOffDays)
	startTime := ParseTime(body.StartTime)
	endTime := ParseTime(body.EndTime)

	tournament := models.Tournament{
		Description:      body.Description,
		StartDate:        startDate,
		EndDate:          endDate,
		StartTime:        startTime,
		EndTime:          endTime,
		HostClubID:       clubID,
		RoundrobinDays:   roundrobinDays,
		RoundrobinCourts: roundRobinCourts,
		PlayOffDays:      playoffDays,
		PlayoffCourts:    playoffCourts,
		GameDuration:     gameDuration,
		Active:           true,
	}
	//fmt.Println(tournament)
	result := initializers.DB.Create(&tournament)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear torneo... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
