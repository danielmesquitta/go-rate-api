package handler

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/go-rate-api/model"
	"github.com/danielmesquitta/go-rate-api/util"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"min=8"`
}

// @BasePath /api/v1
// @Summary Update user
// @Description Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User Identification"
// @Param user body UpdateUserRequest true "User data to Update"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [put]
func UpdateUserHandler(ctx *gin.Context) {
	// Get id from query and validate
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	dto := UpdateUserRequest{}

	ctx.BindJSON(&dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(ctx, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	user := model.User{}

	// Find user
	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	user.Name = dto.Name
	user.Email = dto.Email
	user.Password = dto.Password

	// Save opening
	if err := db.Save(&user).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
