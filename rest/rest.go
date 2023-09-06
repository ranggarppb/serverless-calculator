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

		h.handleSuccess(w, types.CalculatorHistory{Result: history})

		return

	case http.MethodPost:
		var calculator types.CalculatorInput
		if err := json.NewDecoder(r.Body).Decode(&calculator); err != nil || calculator.Input == "" {
			h.handleError(w, types.ErrInvalidInput)
			return
		}

		result, err := h.calculatorService.Calculate(calculator.Input)

		if err != nil {
			h.handleError(w, err)
			return
		}

		h.handleSuccess(w, types.CalculatorResult{Input: calculator.Input, Result: result})

		return

	default:
		h.handleError(w, types.ErrInvalidMethod)

		return
	}
}

func (h *restHandler) handleSuccess(w http.ResponseWriter, result interface{}) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}

func (h *restHandler) handleError(w http.ResponseWriter, err types.WrappedError) {
	w.WriteHeader(err.StatusCode())

	errorResponse := types.ErrorResponse{
		ErrorCode:    err.ErrCode(),
		ErrorMessage: err.Error(),
	}

	json.NewEncoder(w).Encode(errorResponse)
}
