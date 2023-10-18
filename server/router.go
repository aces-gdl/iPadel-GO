package server

import (
	"iPadel-GO/controllers"
	"iPadel-GO/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/v1/security/signup", controllers.Signup)
	router.POST("/v1/security/login", controllers.Login)
	router.GET("/v1/security/validate", middleware.RequireAuth, controllers.Validate)

	router.GET("/v1/catalogs/users", middleware.RequireAuth, controllers.GetUsers)

	router.POST("/v1/catalogs/club", middleware.RequireAuth, controllers.PostClub)
	router.GET("/v1/catalogs/clubs", middleware.RequireAuth, controllers.GetClubs)

	router.POST("/v1/catalogs/category", middleware.RequireAuth, controllers.PostCategory)
	router.GET("/v1/catalogs/categories", middleware.RequireAuth, controllers.GetCatgories)

	router.GET("/v1/catalogs/permissions", middleware.RequireAuth, controllers.GetPermissions)

	router.POST("/v1/catalogs/court", middleware.RequireAuth, controllers.PostCourts)
	router.GET("/v1/catalogs/court", middleware.RequireAuth, controllers.GetCourts)
	router.GET("/v1/catalogs/court/byclub", middleware.RequireAuth, controllers.GetCourtsByClub)
	router.GET("/v1/catalogs/tournaments", middleware.RequireAuth, controllers.GetTournaments)
	router.POST("/v1/catalogs/tournaments", middleware.RequireAuth, controllers.PostTournaments)

	router.POST("/v1/tournament/simulateenrollment", middleware.RequireAuth, controllers.PostSimulateEnrollment)
	router.POST("/v1/tournament/creategroups", middleware.RequireAuth, controllers.PostCreateGroups)
	router.GET("/v1/tournament/getteams", middleware.RequireAuth, controllers.GetEnrolledTeams)
	router.GET("/v1/tournament/getteamsbygroup", middleware.RequireAuth, controllers.GetGroups)
	router.POST("/v1/tournament/createtimeslots", middleware.RequireAuth, controllers.PostCreateTimeSlots)
	router.GET("/v1/tournament/gettimeslots", middleware.RequireAuth, controllers.GetTimeSlots)
	router.GET("/v1/tournament/enrolledteams", middleware.RequireAuth, controllers.GetEnrolledTeams)
	router.GET("/v1/tournament/listgames", middleware.RequireAuth, controllers.GetGames)

	router.POST("/v1/utility/loadusers", middleware.RequireAuth, controllers.PostLoadUsers)

	return router
}
