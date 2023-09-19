package calculation

import (
	"context"
	"net/http"

	"github.com/ranggarppb/serverless-calculator/errors"
)

type ICalculationInput interface {
	GetInput() string
	Validate() errors.WrappedError
	Calculate() string
}

type ICalculatorService interface {
	GetCalculationHistory(context.Context) CalculationHistory
	Calculate(context.Context, ICalculationInput) (CalculationResult, errors.WrappedError)
}

type ICalculatorRestHandler interface {
	HandleReadinessLiveness(context.Context, http.ResponseWriter, *http.Request)
	HandleCalculation(context.Context, http.ResponseWriter, *http.Request)
}
