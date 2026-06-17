package returnrate

import "github.com/IamSBStakumi/asset-simulator/internal/modules/returnrate/internal/domain"

type Response struct {
	MonthlyYieldRate float64 `json:"monthly_yield_rate"`
	AnnualYieldRate  float64 `json:"annual_yield_rate"`
}

func NewResponse(result domain.CalculationResult) Response {
	return Response{
		MonthlyYieldRate: result.MonthlyYieldPercentage,
		AnnualYieldRate:  result.AnnualYieldPercentage,
	}
}
