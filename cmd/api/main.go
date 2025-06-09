package main

import (
	// _ "database/sql"

	// "fmt"
	"errors"
	"log"
	"net/http"

	"github.com/kkhuzzyatov/GameAnalytics/pkg/handlers"
	"github.com/kkhuzzyatov/GameAnalytics/pkg/repository"
	"github.com/kkhuzzyatov/GameAnalytics/pkg/service"
	"github.com/sirupsen/logrus"

	// "github.com/gin-gonic/gin"
	"github.com/kkhuzzyatov/GameAnalytics/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// inits db and returns a *sqlx.DB instance
	// pings the db to check if it is available
	db, err := repository.NewPostgresDB(cfg.PG.DSN)
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}
	defer db.Close()

	// initialize repository, service, and handlers
	// we pass the db instance to the repository constructor
	// and then pass the repository to the service constructor
	// finally, we pass the service to the handlers constructor
	// this way we can use the repository methods in the service methods
	// and the service methods in the handlers methods
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	if errSrv := http.ListenAndServe(cfg.HTTP.Port, handlers.InitRoutes()); errSrv != nil && !errors.Is(errSrv, http.ErrServerClosed) {
		logrus.Fatalf("error starting server: %s", err.Error())
	}

	
	// db, err := sql.Open("postgres", cfg.PG.DSN)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	// r := gin.Default()

	// r.GET("/health", func(c *gin.Context) {
	// 	if err := db.Ping(); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"status": "unhealthy",
	// 		})
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "healthy",
	// 	})
	// })

	// r.Run(fmt.Sprintf("localhost:%s", cfg.HTTP.Port))
}
