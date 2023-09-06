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
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		history := h.calculatorService.GetCalculationHistory()

		calculatorHistory := types.CalculatorHistory{
			Result: history,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(calculatorHistory)

		return

	case http.MethodPost:
		var calculator types.CalculatorInput
		if err := json.NewDecoder(r.Body).Decode(&calculator); err != nil || calculator.Input == "" {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, fmt.Sprintf("error reading request body, %v", err), http.StatusBadRequest)

			return
		}

		result, err := h.calculatorService.Calculate(calculator.Input)

		if err != nil {
			return
		}

		calculatorResult := types.CalculatorResult{
			Result: result,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(calculatorResult)

		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "error method not allowed", http.StatusMethodNotAllowed)

		return
	}
}
