package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kkhuzzyatov/GameAnalytics/pkg/entities"
)

func GetProjectMetric(c *gin.Context) {
	// here we can use `func getUserId(c) (int, error)` to get the user ID from context instead
	// developerId, err := getUserId(c)

	developerId, err := strconv.Atoi(c.Query("id"))
	_ = developerId // stub for now, to avoid unused variable error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var projects []entities.Project = nil // todo: we need to get all projects whose developer has id = developerId
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}
