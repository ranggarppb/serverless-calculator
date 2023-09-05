package calculator

import (
	"fmt"
	"html"
	"net/http"
)

type calculatorService struct{}

func NewCalculatorService() *calculatorService {
	return &calculatorService{}
}

func (c *calculatorService) GetCalculationHistory(w http.ResponseWriter) {
	fmt.Fprint(w, "Hello, World!")
}

func (c *calculatorService) Calculate(w http.ResponseWriter, input string) {

	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(input))
}
