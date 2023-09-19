package rest_test

import (
	"context"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/internal/rest"
	"github.com/ranggarppb/serverless-calculator/mocks"
	c "github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandleCalculation(t *testing.T) {
	testCases := []struct {
		Desc               string
		Ctx                context.Context
		CalculatorService  *mocks.ICalculatorService
		ReqBody            io.Reader
		ExpectedRespBody   string
		ExpectedStatusCode int
	}{
		{
			Desc: "success-processing-request-with-multiple-inputs",
			Ctx:  context.TODO(),
			CalculatorService: func() *mocks.ICalculatorService {
				mockSvc := new(mocks.ICalculatorService)
				input := c.NewCalculationWithMultipleInput([]string{"1", "add", "3", "multiply", "2"})
				mockSvc.On("Calculate", mock.Anything, input).Return(c.CalculationResult{Input: "1 add 3 multiply 2", Result: "7"}, nil)
				return mockSvc
			}(),
			ReqBody: strings.NewReader(
				`{
					"input": "1 add 3 multiply 2"
				}`),
			ExpectedRespBody:   `{"input":"1 add 3 multiply 2","result":"7"}`,
			ExpectedStatusCode: 200,
		},
		{
			Desc: "success-processing-request-with-one-input",
			Ctx:  context.TODO(),
			CalculatorService: func() *mocks.ICalculatorService {
				mockSvc := new(mocks.ICalculatorService)
				input := c.NewCalculationWithOneInput("sqr", []string{"2"})
				mockSvc.On("Calculate", mock.Anything, input).Return(c.CalculationResult{Input: "sqr 2", Result: "4"}, nil)
				return mockSvc
			}(),
			ReqBody: strings.NewReader(
				`{
					"input": "sqr 2"
				}`),
			ExpectedRespBody:   `{"input":"sqr 2","result":"4"}`,
			ExpectedStatusCode: 200,
		},
		{
			Desc: "success-returning-error",
			Ctx:  context.TODO(),
			CalculatorService: func() *mocks.ICalculatorService {
				mockSvc := new(mocks.ICalculatorService)
				mockSvc.On("Calculate", mock.Anything, mock.AnythingOfType("string")).Return("", errors.ErrInvalidInput)
				return mockSvc
			}(),
			ReqBody: strings.NewReader(
				`{
					"input": ""
				}`),
			ExpectedRespBody:   `{"error_code":"INVALID_INPUT","error_message":"Please specify input field"}`,
			ExpectedStatusCode: 400,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.Desc, func(t *testing.T) {
			restHandler := rest.NewCalculatorRestHandler(tC.CalculatorService)
			req := httptest.NewRequest("POST", "/calculation", tC.ReqBody)
			rec := httptest.NewRecorder()
			restHandler.HandleCalculation(tC.Ctx, rec, req)
			require.Equal(t, tC.ExpectedStatusCode, rec.Code)
			require.Equal(t, fmt.Sprintf("%s\n", tC.ExpectedRespBody), rec.Body.String())
		})
	}
}
