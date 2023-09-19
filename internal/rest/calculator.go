package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	c "github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/ranggarppb/serverless-calculator/utils"
)

type restHandler struct {
	calculatorService c.ICalculatorService
}

func NewCalculatorRestHandler(c c.ICalculatorService) *restHandler {
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
		var calculationInput c.CalculationInput
		if err := json.NewDecoder(r.Body).Decode(&calculationInput); err != nil || calculationInput.Input == "" {
			h.handleError(&w, errors.ErrInvalidInput)
			return
		}

		if splittedInput := strings.Split(calculationInput.Input, " "); len(splittedInput) > 0 &&
			(utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, splittedInput[0])) {
			calculationInput.Input = fmt.Sprintf(" %s", calculationInput.Input)
		}

		parsedCalculationInput, err := calculationInput.ParseCalculationInput(c.CalculationHistory{})

		if err != nil {
			h.handleError(&w, err)
		}

		result, err := h.calculatorService.Calculate(ctx, parsedCalculationInput)

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
