package rest_test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ranggarppb/serverless-calculator/mocks"
	"github.com/ranggarppb/serverless-calculator/rest"
	"github.com/ranggarppb/serverless-calculator/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandleCalculation(t *testing.T) {
	testCases := []struct {
		Desc               string
		CalculatorService  *mocks.ICalculatorService
		ReqBody            io.Reader
		ExpectedStatusCode int
	}{
		{
			Desc: "success-processing-normal-request",
			CalculatorService: func() *mocks.ICalculatorService {
				mockSvc := new(mocks.ICalculatorService)
				mockSvc.On("Calculate", mock.AnythingOfType("string")).Return("4", nil)
				return mockSvc
			}(),
			ReqBody: strings.NewReader(
				`{
					"input": "1 add 3"
				}`),
			ExpectedStatusCode: 200,
		},
		{
			Desc: "success-returning-error",
			CalculatorService: func() *mocks.ICalculatorService {
				mockSvc := new(mocks.ICalculatorService)
				mockSvc.On("Calculate", mock.AnythingOfType("string")).Return("", types.ErrInvalidInput)
				return mockSvc
			}(),
			ReqBody: strings.NewReader(
				`{
					"input": ""
				}`),
			ExpectedStatusCode: 400,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.Desc, func(t *testing.T) {
			restHandler := rest.NewRestHandler(tC.CalculatorService)
			req := httptest.NewRequest("POST", "/calculation", tC.ReqBody)
			rec := httptest.NewRecorder()
			restHandler.HandleCalculation(rec, req)
			require.Equal(t, tC.ExpectedStatusCode, rec.Code)
		})
	}
}
