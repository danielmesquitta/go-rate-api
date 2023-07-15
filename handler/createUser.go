package handler

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/go-rate-api/model"
	"github.com/danielmesquitta/go-rate-api/util"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
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
func CreateUserHandler(ctx *gin.Context) {
	dto := CreateUserRequest{}

	ctx.ShouldBindJSON(&dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(ctx, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	user := model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	// Create user
	if err := db.Create(&user).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to create user")
		return
	}

	ctx.Writer.WriteHeader(http.StatusCreated)
}
