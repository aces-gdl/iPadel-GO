package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Correo y/o contraseña son requeridos...",
		})
		return
	}

	user := models.User{
		Email:        body.Email,
		GivenName:    body.GivenName,
		FamilyName:   body.FamilyName,
		GoogleID:     body.GoogleID,
		ImageURL:     body.ImageURL,
		Name:         body.Name,
		PermissionID: body.PermissionID,
		CategoryID:   body.CategoryID,
	}

	if body.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Fallo al convertir password a hash...",
			})
			return
		}
		user.Password = string(hash)
	} else {
		if body.GoogleID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Correo y/o contraseña son requeridos...",
			})
			return
		}
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Correo o clave invalido ...",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email= ?", body.Email)
	if user.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Correo o clave invalido ...",
		})
		return
	}

	if body.GoogleID != "" {
		if body.GoogleID != user.GoogleID {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Cuenta google invalida, Hacer registro primero ...",
			})
			return
		}
	}

	if body.Password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Correo o clave invalido ...",
			})
			return
		}

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear token ...",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*2, "/", "localhost", false, true)
	//
	// respond with token
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"messaje": user,
	})
}
