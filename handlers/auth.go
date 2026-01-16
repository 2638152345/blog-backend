package handlers

import (
	"net/http"
	"time"

	"blog-backend/config"
	"blog-backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		RespondError(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		RespondError(c, http.StatusInternalServerError, "password encrypt failed", err)
		return
	}

	user.Password = string(hashedPassword)
	if err := config.DB.Create(&user).Error; err != nil {
		RespondError(c, http.StatusInternalServerError, "create user failed", err)
		return
	}

	RespondSuccess(c, "register success", gin.H{"user_id": user.ID})
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		RespondError(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	var storedUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		RespondError(c, http.StatusUnauthorized, "invalid username or password", err)
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(storedUser.Password),
		[]byte(user.Password),
	); err != nil {
		RespondError(c, http.StatusUnauthorized, "invalid username or password", err)
		return
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       storedUser.ID,
			"username": storedUser.Username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		},
	)
	tokenString, err := token.SignedString(
		[]byte("your_secret_key"),
	)

	if err != nil {
		RespondError(c, http.StatusInternalServerError, "failed to generate token", err)
		return
	}

	RespondSuccess(c, "login success", gin.H{"token": tokenString})
}
