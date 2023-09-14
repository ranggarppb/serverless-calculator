package rest

import (
	"context"
	"net/http"
)

type ICalculatorRestHandler interface {
	HandleReadinessLiveness(context.Context, http.ResponseWriter, *http.Request)
	HandleCalculation(context.Context, http.ResponseWriter, *http.Request)
}
