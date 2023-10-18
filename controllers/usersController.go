package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")
	var categoryID = c.DefaultQuery("CategoryID", "")
	var PermissionID = c.DefaultQuery("PermissionID", "")
	var ID = c.DefaultQuery("ID", "")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var whereClause = " 1 = 1 "
	if ID != "" {
		whereClause = whereClause + " AND ID = '" + ID + "' "
	}
	if categoryID != "" {
		whereClause = whereClause + " AND category_id = '" + categoryID + "' "
	}
	if PermissionID != "" {
		whereClause = whereClause + " AND permission_id = '" + PermissionID + "'"
	}
	var users []models.User
	results := initializers.DB.Where(whereClause).Limit(intLimit).Offset(offset).Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}

func UpdateUsers(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	user := models.User{
		ID:           body.ID,
		Email:        body.Email,
		GivenName:    body.GivenName,
		FamilyName:   body.FamilyName,
		GoogleID:     body.GoogleID,
		ImageURL:     body.ImageURL,
		Name:         body.Name,
		PermissionID: body.PermissionID,
		CategoryID:   body.CategoryID,
	}

	results := initializers.DB.Save(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": results.RowsAffected, "data": user})

}
