package main

import (
	"avito-intern-backend/internal/database"
	"avito-intern-backend/internal/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()
	r.GET("/api/ping", handlers.PingHandler)
	r.GET("/api/tenders", handlers.GetTendersHandler)
	r.GET("/api/tenders/my", handlers.GetUserTendersHandler)
	r.GET("/api/tenders/:tenderId/status", handlers.GetTenderStatusHandler)
	r.PUT("/api/tenders/:tenderId/status", handlers.UpdateTenderStatusHandler)
	r.PUT("/api/tenders/:tenderId/rollback/:version", handlers.RollbackTenderHandler)
	r.POST("/api/tenders/new", handlers.CreateTenderHandler)
	r.PATCH("/api/tenders/:tenderId/edit", handlers.EditTenderHandler)

	err = r.Run(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
