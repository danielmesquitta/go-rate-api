package handler

import (
	"net/http"
	"time"

	"github.com/danielmesquitta/go-rate-api/model"
	"github.com/danielmesquitta/go-rate-api/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// @BasePath /api/v1
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Request body"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	dto := LoginRequest{}

	// Bind request body to dto
	c.ShouldBindJSON(&dto)

	// Trim and lowercase email
	util.FormatEmail(&dto.Email)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(c, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	user := model.User{}

	// Find user by email
	if err := db.First(&user, "email = ?", dto.Email).Error; err != nil {
		sendError(c, http.StatusNotFound, "user not found")
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		sendError(c, http.StatusUnauthorized, "invalid password")
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	token, err := util.GenerateJwt(user.ID, time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		sendError(c, http.StatusInternalServerError, "error generating token")
		return
	}

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
