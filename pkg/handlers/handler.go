package handlers

import (
	"github.com/kkhuzzyatov/GameAnalytics/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// InitRoutes initializes the routes for the API returning a gin.Engine instance.
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// usersApi := router.Group("/users", h.userIdentity) // we user middleware to check if user is authenticated
	// {

	// }

	return router
}
