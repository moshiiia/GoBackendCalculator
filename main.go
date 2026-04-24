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

func patchCalculations(c echo.Context) error {
	id := c.Param("id")

	var req CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	result, err := CalculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	for i, Calculation := range Calculations {
		if Calculation.ID == id {
			Calculations[i].Expression = req.Expression
			Calculations[i].Result = result
			return c.JSON(http.StatusOK, Calculations[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Calculation not found!"})
}

func deleteCalculations(c echo.Context) error {
	id := c.Param("id")

	for i, calc := range Calculations {
		if calc.ID == id {
			Calculations = append(Calculations[:i], Calculations[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Deleted"})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Calculation not found!"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS()) //запрещает фронтенду ходить в другой backend без разрешения
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculation)
	e.POST("/calculations", postCalculations)

	e.PATCH("/calculations/:id", patchCalculations)
	e.DELETE("/calculations/:id", deleteCalculations)

	err := e.Start("localhost:8080")
	if err != nil {
		return
	}
}
