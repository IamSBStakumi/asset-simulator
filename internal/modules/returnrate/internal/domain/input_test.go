package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculationInputValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       CalculationInput
		expectedErr string
	}{
		{
			name: "valid with initial principal and monthly contribution",
			input: CalculationInput{
				PrincipalAmount:     2_200_000,
				CurrentProfit:       100_000,
				AccumulatedMonths:   12,
				MonthlyContribution: 100_000,
			},
		},
		{
			name: "valid with initial principal only",
			input: CalculationInput{
				PrincipalAmount:     1_000_000,
				CurrentProfit:       0,
				AccumulatedMonths:   12,
				MonthlyContribution: 0,
			},
		},
		{
			name: "negative principal amount",
			input: CalculationInput{
				PrincipalAmount:     -1,
				CurrentProfit:       0,
				AccumulatedMonths:   12,
				MonthlyContribution: 0,
			},
			expectedErr: "元本は0以上で入力してください",
		},
		{
			name: "negative current amount",
			input: CalculationInput{
				PrincipalAmount:     100,
				CurrentProfit:       -101,
				AccumulatedMonths:   12,
				MonthlyContribution: 0,
			},
			expectedErr: "現在資産額は0以上で入力してください",
		},
		{
			name: "zero accumulated months",
			input: CalculationInput{
				PrincipalAmount:     100,
				CurrentProfit:       0,
				AccumulatedMonths:   0,
				MonthlyContribution: 0,
			},
			expectedErr: "積立月数は1以上で入力してください",
		},
		{
			name: "negative monthly contribution",
			input: CalculationInput{
				PrincipalAmount:     100,
				CurrentProfit:       0,
				AccumulatedMonths:   12,
				MonthlyContribution: -1,
			},
			expectedErr: "毎月の積立額は0以上で入力してください",
		},
		{
			name: "principal less than total contribution",
			input: CalculationInput{
				PrincipalAmount:     1_199_999,
				CurrentProfit:       0,
				AccumulatedMonths:   12,
				MonthlyContribution: 100_000,
			},
			expectedErr: "元本は毎月の積立額と積立月数の合計以上で入力してください",
		},
		{
			name: "no initial principal and no monthly contribution",
			input: CalculationInput{
				PrincipalAmount:     0,
				CurrentProfit:       0,
				AccumulatedMonths:   12,
				MonthlyContribution: 0,
			},
			expectedErr: "初期元本または毎月の積立額を入力してください",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.input.Validate()
			if tt.expectedErr == "" {
				require.NoError(t, err)
				return
			}

			require.EqualError(t, err, tt.expectedErr)
		})
	}
}

func TestCalculationInputCurrentAmount(t *testing.T) {
	t.Parallel()

	input := CalculationInput{
		PrincipalAmount: 1_000_000,
		CurrentProfit:   250_000,
	}

	require.Equal(t, int64(1_250_000), input.CurrentAmount())
}

func TestCalculationInputInitialPrincipal(t *testing.T) {
	t.Parallel()

	input := CalculationInput{
		PrincipalAmount:     2_200_000,
		AccumulatedMonths:   12,
		MonthlyContribution: 100_000,
	}

	require.Equal(t, int64(1_000_000), input.InitialPrincipal())
}
