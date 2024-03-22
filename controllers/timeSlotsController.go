package controllers

import (
	"fmt"
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostCreateTimeSlots(c *gin.Context) {

	TournamentID := c.GetUint("TournamentID")

	var tournament models.Tournament
	initializers.DB.Where("id = ?", TournamentID).First(&tournament)

	var myLocation, _ = time.LoadLocation("America/Mexico_City")
	var dailyFirstGameStart = time.Date(tournament.StartDate.Year(), tournament.StartDate.Month(), tournament.StartDate.Day(), tournament.StartTime.Hour(), tournament.StartTime.Minute(), 0, 0, myLocation)
	var dailyLastGameStart = time.Date(tournament.EndDate.Year(), tournament.EndDate.Month(), tournament.EndDate.Day(), tournament.EndTime.Hour(), tournament.EndTime.Minute(), 0, 0, myLocation)
	var gameDuration = 60 * time.Minute // in minutes
	var availableCourtCount = tournament.RoundrobinCourts
	var daysForRoundRobin = tournament.RoundrobinDays

	var gamesPerDayPerCourt = int(tournament.EndTime.Sub(tournament.StartTime) / gameDuration)
	var gamesPerDay = gamesPerDayPerCourt * availableCourtCount
	var roundrobinAvailableSlots = gamesPerDay * daysForRoundRobin

	fmt.Println("Games per court per day ", int(dailyLastGameStart.Sub(dailyFirstGameStart)/gameDuration))
	fmt.Println("Games per day", gamesPerDay)
	fmt.Println("Total round robin available slots", roundrobinAvailableSlots)

	runningDate := dailyFirstGameStart
	slotCounter := 1
	for days := 1; days <= daysForRoundRobin; days++ {

		for courtNumber := 1; courtNumber <= availableCourtCount; courtNumber++ {
			tempRunningDate := runningDate
			for i := 0; i < gamesPerDayPerCourt; i++ {
				var timeSlot models.TournamentTimeSlots
				timeSlot.Description = fmt.Sprintf("%02d", slotCounter)
				timeSlot.CourtNumber = courtNumber
				timeSlot.StartTime = runningDate
				timeSlot.EndTime = runningDate.Add(gameDuration)
				timeSlot.TournamentID = TournamentID
				timeSlot.Taken = false
				timeSlot.Active = true
				fmt.Println("Resultado : ", timeSlot)
				initializers.DB.Debug().Create(&timeSlot)
				slotCounter = slotCounter + 1
				runningDate = runningDate.Add(gameDuration)
			}
			runningDate = tempRunningDate
		}
		runningDate = dailyFirstGameStart.AddDate(0, 0, days)
	}

}

func GetTimeSlots(c *gin.Context) {
	TournamentID := c.GetUint("TournamentID")

	//FilterDate, _ := time.ParseInLocation(time.RFC3339, c.DefaultQuery("FilterDate", ""), initializers.DB.NowFunc().Location())
	FilterDate := c.DefaultQuery("FilterDate", "")

	var tournament models.Tournament
	initializers.DB.Where("id = ?", TournamentID).First(&tournament)

	var timeSlotsRecords []models.TournamentTimeSlots

	myQuery := `Select *
				From tournament_time_slots 
		 		where  1 = 1 
					and tournament_id = ` + strconv.FormatUint(uint64(TournamentID), 10) + ` 
					and start_time  between '` + FilterDate + `T00:00:00.000-06:00' and '` + FilterDate + `T23:59:59.000-06:00'
		 		Order By start_time, court_number`

	results := initializers.DB.Debug().Raw(myQuery).Scan(&timeSlotsRecords)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(timeSlotsRecords), "data": timeSlotsRecords})
}
