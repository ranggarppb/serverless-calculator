package types

import (
	"net/http"
)

type CalculatorInput struct {
	Input string `json:"input"`
}

type ICalculatorService interface {
	GetCalculationHistory(w http.ResponseWriter)
	Calculate(w http.ResponseWriter, input string)
}
