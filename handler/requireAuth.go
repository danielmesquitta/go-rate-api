package handler

import (
	"net/http"
	"time"

	"github.com/danielmesquitta/go-rate-api/model"
	"github.com/danielmesquitta/go-rate-api/util"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	// Get the JWT token from the cookie
	bearerToken := c.Request.Header.Get("Authorization")

	// Check if the token is present
	if bearerToken == "" {
		sendError(c, http.StatusUnauthorized, "missing authorization token")
		return
	}

	// Format the token string
	tokenString := bearerToken[7:]

	// Parse and validate the JWT token
	claims, err := util.ParseAndValidateJwt(tokenString)
	if err != nil {
		sendError(c, http.StatusUnauthorized, "failed to validate authorization cookie")
		return
	}

	// Check the expiration time
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		sendError(c, http.StatusUnauthorized, "token expired")
		return
	}

	// Find user by ID
	user := model.User{}
	if err := db.First(&user, claims["sub"]).Error; err != nil {
		sendError(c, http.StatusNotFound, "user not exists")
		return
	}

	// Set user in context
	c.Set("user", user)

	// Call next handler
	c.Next()
}
