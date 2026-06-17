package returnrate

import "github.com/IamSBStakumi/asset-simulator/internal/modules/returnrate/internal/domain"

type Request struct {
	PrincipalAmount     int64 `json:"principal_amount"`
	CurrentProfit       int64 `json:"current_profit"`
	AccumulatedMonths   int   `json:"accumulated_months"`
	MonthlyContribution int64 `json:"monthly_contribution"`
}

func (r Request) ToInput() domain.CalculationInput {
	return domain.CalculationInput{
		PrincipalAmount:     r.PrincipalAmount,
		CurrentProfit:       r.CurrentProfit,
		AccumulatedMonths:   r.AccumulatedMonths,
		MonthlyContribution: r.MonthlyContribution,
	}
}
