package domain

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculatorCalculate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		input               CalculationInput
		expectedMonthlyRate float64
		expectedAnnualRate  float64
	}{
		{
			name: "zero yield when current amount equals total principal",
			input: CalculationInput{
				PrincipalAmount:     2_200_000,
				CurrentProfit:       0,
				AccumulatedMonths:   12,
				MonthlyContribution: 100_000,
			},
			expectedMonthlyRate: 0,
			expectedAnnualRate:  0,
		},
		{
			name: "positive yield",
			input: inputForMonthlyRate(
				1_000_000,
				100_000,
				12,
				0.01,
			),
			expectedMonthlyRate: 1,
			expectedAnnualRate:  (math.Pow(1.01, 12) - 1) * 100,
		},
		{
			name: "negative yield",
			input: inputForMonthlyRate(
				1_000_000,
				100_000,
				12,
				-0.01,
			),
			expectedMonthlyRate: -1,
			expectedAnnualRate:  (math.Pow(0.99, 12) - 1) * 100,
		},
	}

	calculator := NewCalculator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := calculator.Calculate(tt.input)

			require.NoError(t, err)
			require.InDelta(t, tt.expectedMonthlyRate, result.MonthlyYieldPercentage, 0.00001)
			require.InDelta(t, tt.expectedAnnualRate, result.AnnualYieldPercentage, 0.0001)
		})
	}
}

func TestCalculatorCalculateReturnsErrorWhenCurrentAmountIsTooSmall(t *testing.T) {
	t.Parallel()

	calculator := NewCalculator()
	input := CalculationInput{
		PrincipalAmount:     1_000_000,
		CurrentProfit:       -1_000_000,
		AccumulatedMonths:   12,
		MonthlyContribution: 0,
	}

	_, err := calculator.Calculate(input)

	require.EqualError(t, err, "現在資産額が小さすぎるため利回りを計算できません")
}

func inputForMonthlyRate(
	initialPrincipal int64,
	monthlyContribution int64,
	accumulatedMonths int,
	monthlyRate float64,
) CalculationInput {
	principalAmount := initialPrincipal + monthlyContribution*int64(accumulatedMonths)
	currentAmount := futureValue(
		initialPrincipal,
		monthlyContribution,
		accumulatedMonths,
		monthlyRate,
	)

	return CalculationInput{
		PrincipalAmount:     principalAmount,
		CurrentProfit:       int64(math.Round(currentAmount)) - principalAmount,
		AccumulatedMonths:   accumulatedMonths,
		MonthlyContribution: monthlyContribution,
	}
}

func futureValue(
	initialPrincipal int64,
	monthlyContribution int64,
	accumulatedMonths int,
	monthlyRate float64,
) float64 {
	amount := float64(initialPrincipal)
	for range accumulatedMonths {
		amount = amount*(1+monthlyRate) + float64(monthlyContribution)
	}

	return amount
}
