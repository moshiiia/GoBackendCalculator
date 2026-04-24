package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getCalculation(c echo.Context) error {
	return c.JSON(http.StatusOK, Calculations)
}

func postCalculations(c echo.Context) error {
	var req CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	result, err := CalculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: req.Expression,
		Result:     result,
	}
	Calculations = append(Calculations, calc)
	return c.JSON(http.StatusCreated, calc)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS()) //запрещает фронтенду ходить в другой backend без разрешения
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculation)
	e.POST("/calculations", postCalculations)

	err := e.Start("localhost:8080")
	if err != nil {
		return
	}
}
