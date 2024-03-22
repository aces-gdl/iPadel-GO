package controllers

import (
	"bufio"
	"fmt"
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostLoadUsers(c *gin.Context) {
	const PWD_DEFAULT = "APJ123"

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Archivo no esta presente...",
		})
		return
	}

	CategoryID := c.PostForm("CategoryID")
	PermissionID := c.PostForm("PermissionID")

	var Category models.Category
	result := initializers.DB.Where("ID = ?", CategoryID).First(&Category)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error buscando Categoria...",
		})
		return
	}

	var Permission models.Permission
	result = initializers.DB.Where("ID = ?", PermissionID).First(&Permission)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error buscando Permisos...",
		})
		return
	}

	fileToImport, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Archivo no esta presente...",
		})
		return
	}
	defer fileToImport.Close()

	fileScanner := bufio.NewScanner(fileToImport)

	fileScanner.Split(bufio.ScanLines)

	// Create password default

	hash, err := bcrypt.GenerateFromPassword([]byte(PWD_DEFAULT), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al convertir password a hash...",
		})
		return
	}

	var pwdDefault = string(hash)
	for fileScanner.Scan() {
		var user models.User
		arrayUser := strings.Split(fileScanner.Text(), ",")
		user.GivenName = arrayUser[0]
		user.FamilyName = arrayUser[1]
		user.Email = arrayUser[2]
		user.Ranking, _ = strconv.Atoi(arrayUser[3])
		user.Name = fmt.Sprintf("%s, %s", user.GivenName, user.FamilyName)
		user.CategoryID = Category.ID
		user.PermissionID = Permission.ID
		user.Password = pwdDefault

		initializers.DB.Create(&user)
	}

	c.JSON(http.StatusOK, gin.H{})
}
