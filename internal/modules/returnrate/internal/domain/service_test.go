package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculationServiceCalculate(t *testing.T) {
	t.Parallel()

	service := NewCalculationService(*NewCalculator())
	input := CalculationInput{
		PrincipalAmount:     1_000_000,
		CurrentProfit:       0,
		AccumulatedMonths:   12,
		MonthlyContribution: 0,
	}

	result, err := service.Calculate(input)

	require.NoError(t, err)
	require.Equal(t, CalculationResult{}, result)
}

func TestCalculationServiceCalculateReturnsValidationError(t *testing.T) {
	t.Parallel()

	service := NewCalculationService(*NewCalculator())
	input := CalculationInput{
		PrincipalAmount:     0,
		CurrentProfit:       0,
		AccumulatedMonths:   12,
		MonthlyContribution: 0,
	}

	_, err := service.Calculate(input)

	require.EqualError(t, err, "初期元本または毎月の積立額を入力してください")
}
