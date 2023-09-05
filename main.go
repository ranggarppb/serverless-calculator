package serverless_calculator

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/ranggarppb/serverless-calculator/rest"
)

func init() {
	functions.HTTP("Calculate", Calculate)
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GetMethod:
		rest.HandleGet(w, r)
	case PostMethod:
		rest.HandlePost(w, r)
	}

}
