package calculator

import (
	"context"
	"net/http"

	"github.com/ranggarppb/serverless-calculator/errors"
)

type ICalculationInput interface {
	Validate() errors.WrappedError
}

type ICalculatorService interface {
	GetCalculationHistory(context.Context) CalculationHistory
	Calculate(context.Context, string) (CalculationResult, errors.WrappedError)
}

type ICalculatorRestHandler interface {
	HandleReadinessLiveness(context.Context, http.ResponseWriter, *http.Request)
	HandleCalculation(context.Context, http.ResponseWriter, *http.Request)
}
