package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {
	var categoryID = c.DefaultQuery("CategoryID", "")
	var PermissionID = c.DefaultQuery("PermissionID", "")
	var ID = c.DefaultQuery("ID", "")
	var SearchString = c.DefaultQuery("SearchString", "")

	var whereClause = " where 1 = 1 "
	if ID != "" {
		whereClause = whereClause + " AND u.ID = '" + ID + "' "
	}
	if categoryID != "" {
		whereClause = whereClause + " AND u.category_id = '" + categoryID + "' "
	}
	if PermissionID != "" {
		whereClause = whereClause + " AND u.permission_id = '" + PermissionID + "'"
	}
	if SearchString != "" {
		whereClause = whereClause + " AND u.name ilike  '%" + SearchString + "%'"
	}

	queryString := `SELECT u.*, c.description as category_description, p.description as permission_description 
				 FROM "users" u
					inner join categories c on u.category_id = c.id 
					inner join permissions p on u.permission_id = p.id ` + whereClause
	type userExtended struct {
		ID                    uuid.UUID
		CategoryID            uuid.UUID
		PermissionIDID        uuid.UUID
		Name                  string
		FamilyName            string
		GivenName             string
		Email                 string
		CategoryDescription   string
		PermissionID          uuid.UUID
		MemberSince           time.Time
		Phone                 string
		HasPicture            int
		Birthday              time.Time
		PermissionDescription string
		Ranking               int
	}

	var usersExtended []userExtended

	results := initializers.DB.Debug().Raw(queryString).Where(whereClause).Scan(&usersExtended)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(usersExtended), "data": usersExtended})
}

func UpdateUsers(c *gin.Context) {
	var body struct {
		ID           string
		Email        string
		Name         string
		FamilyName   string
		GivenName    string
		PermissionID string
		Ranking      string
		CategoryID   string
		MemberSince  string
		Birthday     string
		HasPicture   string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	ID, _ := uuid.Parse(body.ID)
	MemberSince, _ := time.ParseInLocation(time.RFC3339, body.MemberSince, initializers.DB.NowFunc().Location())
	Birthday, _ := time.ParseInLocation(time.RFC3339, body.Birthday, initializers.DB.NowFunc().Location())
	PermissionID, _ := uuid.Parse(body.PermissionID)
	CategoryID, _ := uuid.Parse(body.CategoryID)
	HasPicture, _ := strconv.Atoi(body.HasPicture)
	user := models.User{
		ID:           ID,
		Email:        body.Email,
		Name:         body.Name,
		FamilyName:   body.FamilyName,
		GivenName:    body.GivenName,
		PermissionID: PermissionID,
		Ranking:      0,
		CategoryID:   CategoryID,
		MemberSince:  MemberSince,
		Birthday:     Birthday,
		HasPicture:   HasPicture,
	}
	results := initializers.DB.Save(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": results.RowsAffected, "data": user})

}

func PostUsers(c *gin.Context) {
	var body struct {
		Email        string
		Name         string
		FamilyName   string
		GivenName    string
		PermissionID string
		Ranking      string
		CategoryID   string
		MemberSince  string
		Birthday     string
		HasPicture   string
		Phone        string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	MemberSince, _ := time.ParseInLocation(time.RFC3339, body.MemberSince, initializers.DB.NowFunc().Location())
	Birthday, _ := time.ParseInLocation(time.RFC3339, body.Birthday, initializers.DB.NowFunc().Location())
	PermissionID, _ := uuid.Parse(body.PermissionID)
	CategoryID, _ := uuid.Parse(body.CategoryID)
	HasPicture, _ := strconv.Atoi(body.HasPicture)
	user := models.User{
		Email:        body.Email,
		Name:         body.Name,
		FamilyName:   body.FamilyName,
		GivenName:    body.GivenName,
		PermissionID: PermissionID,
		Ranking:      0,
		CategoryID:   CategoryID,
		MemberSince:  MemberSince,
		Birthday:     Birthday,
		Phone:        body.Phone,
		HasPicture:   HasPicture,
	}

	results := initializers.DB.Create(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": user.ID,
	})
}
