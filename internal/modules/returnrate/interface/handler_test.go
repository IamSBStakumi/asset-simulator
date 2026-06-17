package returnrate

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerCalculate(t *testing.T) {
	t.Parallel()

	handler := NewHandler()
	request := Request{
		PrincipalAmount:     2_200_000,
		CurrentProfit:       195_075,
		AccumulatedMonths:   12,
		MonthlyContribution: 100_000,
	}

	response, err := handler.Calculate(request)

	require.NoError(t, err)
	require.InDelta(t, 1, response.MonthlyYieldRate, 0.00001)
	require.InDelta(t, (math.Pow(1.01, 12)-1)*100, response.AnnualYieldRate, 0.0001)
}

func TestHandlerCalculateReturnsError(t *testing.T) {
	t.Parallel()

	handler := NewHandler()
	request := Request{
		PrincipalAmount:     0,
		CurrentProfit:       0,
		AccumulatedMonths:   12,
		MonthlyContribution: 0,
	}

	_, err := handler.Calculate(request)

	require.EqualError(t, err, "初期元本または毎月の積立額を入力してください")
}
