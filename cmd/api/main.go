package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkhuzzyatov/GameAnalytics/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(cfg.PG.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(ctx *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
			})
		}

		if err := sqlDB.Ping(); err != nil {
			ctx.Status(http.StatusNotFound)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "heathy",
		})
	})

	r.Run(fmt.Sprintf("localhost:%s", cfg.HTTP.Port))
}
