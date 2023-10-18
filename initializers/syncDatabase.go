package initializers

import "iPadel-GO/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{},
		&models.Club{},
		&models.Court{},
		&models.Permission{},
		&models.Tournament{},
		&models.Category{},
		&models.Permission{},
		&models.TournamentTeam{},
		&models.TournamentEnrollment{},
		&models.TournamentGroup{},
		&models.TournamentTeamByGroup{},
		&models.TournamentTimeSlots{},
		&models.TournamentCategories{},
		&models.TournamentGames{},
		&models.TournamentTimeSlots{})
}
