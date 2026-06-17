package simulation

import "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/internal/domain"

type Request struct {
	Principal           int64   `json:"principal"`
	CurrentProfit       int64   `json:"current_profit"`
	InvestedYears       int     `json:"invested_years"`
	MonthlyContribution int64   `json:"monthly_contribution"`
	AnnualYieldRate     float64 `json:"annual_yield_rate"`
	TargetYears         []int   `json:"target_years"`
}

func (r Request) ToInput() domain.SimulationInput {
	return domain.SimulationInput{
		PrincipalAmount:       r.Principal,
		CurrentProfit:         r.CurrentProfit,
		AccumulatedYears:      r.InvestedYears,
		MonthlyContribution:   r.MonthlyContribution,
		AnnualYieldPercentage: r.AnnualYieldRate,
		ProjectionYears:       r.TargetYears,
	}
}
