package domain

import "errors"

type SimulationInput struct {
	PrincipalAmount       int64   // 現在の元本
	CurrentProfit         int64   // 現在の利益
	AccumulatedYears      int     // 投資年数
	MonthlyContribution   int64   // 毎月の積立額
	AnnualYieldPercentage float64 // 利回り
	ProjectionYears       []int   // 結果出力対象の年数
}

func DefaultProjectionYears() []int {
	return []int{1, 3, 5, 10, 20, 30}
}

// 未指定の入力値にデフォルト値を補完する
func (input SimulationInput) Normalize() SimulationInput {
	if len(input.ProjectionYears) == 0 {
		input.ProjectionYears = DefaultProjectionYears()
	}

	return input
}

// 入力値の検証
func (input SimulationInput) Validate() error {
	if input.PrincipalAmount < 0 {
		return errors.New("元本は0以上で入力してください")
	}

	if input.CurrentAmount() < 0 {
		return errors.New("現在資産額は0以上で入力してください")
	}

	if input.AccumulatedYears < 0 {
		return errors.New("積立済み年数は0以上で入力してください")
	}

	if input.MonthlyContribution < 0 {
		return errors.New("毎月の積立額は0以上で入力してください")
	}

	if input.AnnualYieldPercentage < 0 {
		return errors.New("年利は0以上で入力してください")
	}

	for _, year := range input.ProjectionYears {
		if year <= 0 {
			return errors.New("出力対象年数は1以上で入力してください")
		}
	}

	return nil
}

// 現在の資産額を返す
func (input SimulationInput) CurrentAmount() int64 {
	return input.PrincipalAmount + input.CurrentProfit
}

// 年利から月利を計算して返す
func (input SimulationInput) MonthlyRate() float64 {
	return input.AnnualYieldPercentage / 100 / 12
}
