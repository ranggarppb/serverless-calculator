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
	switch r.Method {
	case http.MethodOptions:
		h.handlePreflight(w)
		return
	case http.MethodGet:
		fmt.Fprint(w, "Server OK")
		return
	default:
		h.handlePreflight(w)
		return
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

	case http.MethodOptions:
		h.handlePreflight(w)

		return

	default:
		h.handleError(w, types.ErrInvalidMethod)

		return
	}
}

func (h *restHandler) handlePreflight(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")
	w.WriteHeader(http.StatusNoContent)
}

func (h *restHandler) handleSuccess(w http.ResponseWriter, result interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(result)
}

func (h *restHandler) handleError(w http.ResponseWriter, err types.WrappedError) {
	w.WriteHeader(err.StatusCode())
	w.Header().Set("Access-Control-Allow-Origin", "*")

	errorResponse := types.ErrorResponse{
		ErrorCode:    err.ErrCode(),
		ErrorMessage: err.Error(),
	}

	json.NewEncoder(w).Encode(errorResponse)
}
