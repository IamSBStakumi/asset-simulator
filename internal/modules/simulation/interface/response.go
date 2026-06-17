package simulation

import "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/internal/domain"

// Response は外部へ返却するシミュレーション結果です。
type Response struct {
	Results []YearlyResult `json:"results"`
}

// YearlyResult は特定年数後の資産予測結果です。
type YearlyResult struct {
	// 経過年数
	Year int `json:"year"`

	// 予想資産額
	TotalAmount int64 `json:"total_amount"`

	// 元本合計
	PrincipalAmount int64 `json:"principal_amount"`

	// 利益
	ProfitAmount int64 `json:"profit_amount"`
}

// NewResponse はドメイン層のResultから外部返却用Responseを生成します。
func NewResponse(result domain.SimulationResult) Response {
	results := make([]YearlyResult, 0, len(result.Rows))

	for _, yearlyResult := range result.Rows {
		results = append(results, YearlyResult{
			Year:            yearlyResult.Year,
			TotalAmount:     yearlyResult.TotalAmount,
			PrincipalAmount: yearlyResult.Principal,
			ProfitAmount:    yearlyResult.Profit,
		})
	}

	return Response{
		Results: results,
	}
}
