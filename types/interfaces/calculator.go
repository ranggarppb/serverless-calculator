package interfaces

import (
	"context"
	"net/http"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/types/structs"
)

type ICalculatorService interface {
	GetCalculationHistory(context.Context) structs.CalculationHistory
	Calculate(context.Context, string) (structs.CalculationResult, errors.WrappedError)
}

type ICalculatorRestHandler interface {
	HandleReadinessLiveness(context.Context, http.ResponseWriter, *http.Request)
	HandleCalculation(context.Context, http.ResponseWriter, *http.Request)
}
