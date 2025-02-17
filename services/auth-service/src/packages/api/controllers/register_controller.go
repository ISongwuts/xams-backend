package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/internal/models"
	"github.com/xams-backend/services/auth-service/src/packages/auth"
)

type (
	RegisterController struct {
		Auth auth.IAuthentication
	}
)

func (rc *RegisterController) Register(context *gin.Context) {
	var request models.User
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userResponse, token, err := rc.Auth.Register(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	cookieName := os.Getenv("COOKIE_JWT_TOKEN")

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie(cookieName, token, 3600 * 24 * 30, "", "", false, true)

	context.JSON(http.StatusOK, userResponse)
}