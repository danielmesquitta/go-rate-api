package controller

import (
	"net/http"
	"time"

	"github.com/danielmesquitta/go-rate-api/model"
	"github.com/gin-gonic/gin"
)

type ShowUserResponse struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
}

// @BasePath /api/v1
// @Summary Show user
// @Description Show an user
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User identification"
// @Success 200 {object} ShowUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /user [get]
func ShowUserController(ctx *gin.Context) {
	// Get id from query and validate
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	user := model.User{}

	// Find user

	// Join user address
	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	ctx.JSON(http.StatusOK, user)
}
