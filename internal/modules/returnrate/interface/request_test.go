package returnrate

import (
	"testing"

	"github.com/IamSBStakumi/asset-simulator/internal/modules/returnrate/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestRequestToInput(t *testing.T) {
	t.Parallel()

	request := Request{
		PrincipalAmount:     1_000_000,
		CurrentProfit:       250_000,
		AccumulatedMonths:   12,
		MonthlyContribution: 100_000,
	}

	input := request.ToInput()

	require.Equal(t, domain.CalculationInput{
		PrincipalAmount:     1_000_000,
		CurrentProfit:       250_000,
		AccumulatedMonths:   12,
		MonthlyContribution: 100_000,
	}, input)
}
