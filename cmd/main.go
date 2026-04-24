package main

import (
	"CalculatorAppBackend/internal/calculationService"
	"CalculatorAppBackend/internal/db"
	"CalculatorAppBackend/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()

	if err != nil {
		log.Fatalf("Could not connect to DB")
	}

	e := echo.New()
	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS()) //запрещает фронтенду ходить в другой backend без разрешения
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculation)
	e.POST("/calculations", calcHandlers.PostCalculations)

	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
