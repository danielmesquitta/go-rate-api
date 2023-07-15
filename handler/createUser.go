package handler

import (
	"net/http"
	"strings"

	"github.com/danielmesquitta/go-rate-api/model"
	"github.com/danielmesquitta/go-rate-api/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// @BasePath /api/v1
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
func CreateUserHandler(c *gin.Context) {
	dto := CreateUserRequest{}

	// Bind request body to DTO
	c.ShouldBindJSON(&dto)

	// Trim and lowercase email
	util.FormatEmail(&dto.Email)

	// Trim name
	dto.Name = strings.TrimSpace(dto.Name)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(c, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	// Check if user with same email already exists
	if err := db.First(&model.User{}, "email = ?", dto.Email).Error; err == nil {
		sendError(c, http.StatusBadRequest, "user registered with this email already exists")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "failed to create hashed password")
		return
	}

	// Create user model
	user := model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: string(hashedPassword),
	}

	// Create user
	if err := db.Create(&user).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "failed to create user")
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}
