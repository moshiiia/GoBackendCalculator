package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

// метаданные для полей структуры соотношение названий в коде и в JSON`...`
type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

// динамический массив(слайс) структур Calculation
var Calculations = []Calculation{}

func CalculateExpression(expression string) (string, error) {
	//проверяем, что можно привести к мат. выражению
	expr, err := govaluate.NewEvaluableExpression(expression) //:= — объявление + присваивание
	if err != nil {
		return "", err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}
