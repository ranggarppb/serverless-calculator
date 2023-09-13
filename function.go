package serverless_calculator

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/ranggarppb/serverless-calculator/calculator"
	"github.com/ranggarppb/serverless-calculator/internal/function"
	"github.com/ranggarppb/serverless-calculator/internal/rest"
)

func init() {
	calculatorService := calculator.NewCalculatorService()
	calculatorRestHandler := rest.NewCalculatorRestHandler(calculatorService)

	functions.HTTP("Calculate", function.CreateCalculateFunction(calculatorRestHandler))
}
