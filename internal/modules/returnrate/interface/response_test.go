package returnrate

import (
	"testing"

	"github.com/IamSBStakumi/asset-simulator/internal/modules/returnrate/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestNewResponse(t *testing.T) {
	t.Parallel()

	response := NewResponse(domain.CalculationResult{
		MonthlyYieldPercentage: 1.23,
		AnnualYieldPercentage:  15.75,
	})

	require.Equal(t, Response{
		MonthlyYieldRate: 1.23,
		AnnualYieldRate:  15.75,
	}, response)
}
