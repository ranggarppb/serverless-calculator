package function

import (
	"net/http"

	"github.com/ranggarppb/serverless-calculator/types"
)

func CreateCalculateFunction(calculatorRestHandler types.ICalculatorRestHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/calculation":
			calculatorRestHandler.HandleCalculation(w, r)
		case "/":
			calculatorRestHandler.HandleReadinessLiveness(w, r)
		}
	}
}
