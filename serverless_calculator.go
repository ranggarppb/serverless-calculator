package serverless_calculator

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/ranggarppb/serverless-calculator/rest"
	calculator "github.com/ranggarppb/serverless-calculator/service"
)

func init() {
	functions.HTTP("Calculate", Calculate)
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	calculatorService := calculator.NewCalculatorService()

	restHandler := rest.NewRestHandler(calculatorService)

	switch r.URL.Path {
	case "/calculation":
		restHandler.HandleCalculation(w, r)
	case "/":
		restHandler.HandleReadinessLiveness(w, r)
	}
}
