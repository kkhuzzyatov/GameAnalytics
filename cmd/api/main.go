package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkhuzzyatov/GameAnalytics/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.PG.DSN)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "unhealthy",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	r.Run(fmt.Sprintf("localhost:%s", cfg.HTTP.Port))
}
