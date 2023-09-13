package rest

import "net/http"

type ICalculatorRestHandler interface {
	HandleReadinessLiveness(http.ResponseWriter, *http.Request)
	HandleCalculation(http.ResponseWriter, *http.Request)
}
