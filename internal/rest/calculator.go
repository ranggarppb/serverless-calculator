package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/types/interfaces"
	"github.com/ranggarppb/serverless-calculator/types/structs"
)

type restHandler struct {
	calculatorService interfaces.ICalculatorService
}

func NewCalculatorRestHandler(c interfaces.ICalculatorService) *restHandler {
	return &restHandler{
		calculatorService: c,
	}
}

func (h *restHandler) HandleReadinessLiveness(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		h.handlePreflight(&w)
		return
	case http.MethodGet:
		fmt.Fprint(w, "Server OK")
		return
	default:
		h.handlePreflight(&w)
		return
	}
}

func (h *restHandler) HandleCalculation(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodGet:
		history := h.calculatorService.GetCalculationHistory(ctx)

		h.handleSuccess(&w, history)

		return

	case http.MethodPost:
		var calculator structs.CalculationInput
		if err := json.NewDecoder(r.Body).Decode(&calculator); err != nil || calculator.Input == "" {
			h.handleError(&w, errors.ErrInvalidInput)
			return
		}

		result, err := h.calculatorService.Calculate(ctx, calculator.Input)

		if err != nil {
			h.handleError(&w, err)
			return
		}

		h.handleSuccess(&w, result)

		return

	case http.MethodOptions:
		h.handlePreflight(&w)

		return

	default:
		h.handleError(&w, errors.ErrInvalidMethod)

		return
	}
}

func (h *restHandler) handlePreflight(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Max-Age", "3600")
	(*w).WriteHeader(http.StatusNoContent)
}

func (h *restHandler) handleSuccess(w *http.ResponseWriter, result interface{}) {
	(*w).WriteHeader(http.StatusOK)

	json.NewEncoder((*w)).Encode(result)
}

func (h *restHandler) handleError(w *http.ResponseWriter, err errors.WrappedError) {
	(*w).WriteHeader(err.StatusCode())

	errorResponse := errors.ErrorResponse{
		ErrorCode:    err.ErrCode(),
		ErrorMessage: err.Error(),
	}

	json.NewEncoder((*w)).Encode(errorResponse)
}
