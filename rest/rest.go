package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ranggarppb/serverless-calculator/types"
)

type restHandler struct {
	calculatorService types.ICalculatorService
}

func NewRestHandler(c types.ICalculatorService) *restHandler {
	return &restHandler{
		calculatorService: c,
	}
}

func (h *restHandler) HandleReadinessLiveness(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Server OK")
	}
}

func (h *restHandler) HandleCalculation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.calculatorService.GetCalculationHistory(w)
	case http.MethodPost:
		var calculator types.CalculatorInput
		if err := json.NewDecoder(r.Body).Decode(&calculator); err != nil {
			fmt.Fprint(w, "Hello, World!")
			return
		}
		if calculator.Input == "" {
			fmt.Fprint(w, "Hello, World!")
			return
		}
		h.calculatorService.Calculate(w, calculator.Input)
	}
}
