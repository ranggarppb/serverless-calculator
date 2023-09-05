package rest

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/ranggarppb/serverless-calculator/types"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	var calculator types.CalculatorInput
	if err := json.NewDecoder(r.Body).Decode(&calculator); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if calculator.Input == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(calculator.Input))
}
