package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkhuzzyatov/GameAnalytics/pkg/entities"
)

func UpdatePLayer(c *gin.Context) {
	var player entities.Player

	err := c.ShouldBindJSON(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Player is updated",
		"player":  player,
	})

	// we need to update player in db here
}
