package function

import (
	"context"
	"net/http"

	c "github.com/ranggarppb/serverless-calculator/types/calculator"
)

func CreateCalculateFunction(calculatorRestHandler c.ICalculatorRestHandler) func(http.ResponseWriter, *http.Request) {
	ctx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/calculation":
			calculatorRestHandler.HandleCalculation(ctx, w, r)
		case "/":
			calculatorRestHandler.HandleReadinessLiveness(ctx, w, r)
		}
	}
}
