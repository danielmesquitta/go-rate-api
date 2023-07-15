package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ShowMeResponse struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
}

// @BasePath /api/v1
// @Summary Show me
// @Description Show me
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} ShowMeResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /me [get]
func ShowMeHandler(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		sendError(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	c.JSON(http.StatusOK, user)
}
