package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gameAnalytics "github.com/kkhuzzyatov/GameAnalytics"
)

func (h *Handler) signUp(c *gin.Context) {
	var input gameAnalytics.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid login or password")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
