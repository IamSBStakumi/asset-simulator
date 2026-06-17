package domain

import "errors"

type CalculationInput struct {
	PrincipalAmount     int64
	CurrentProfit       int64
	AccumulatedMonths   int
	MonthlyContribution int64
}

func (input CalculationInput) Validate() error {
	if input.PrincipalAmount < 0 {
		return errors.New("元本は0以上で入力してください")
	}

	if input.CurrentAmount() < 0 {
		return errors.New("現在資産額は0以上で入力してください")
	}

	if input.AccumulatedMonths <= 0 {
		return errors.New("積立月数は1以上で入力してください")
	}

	if input.MonthlyContribution < 0 {
		return errors.New("毎月の積立額は0以上で入力してください")
	}

	if input.InitialPrincipal() < 0 {
		return errors.New("元本は毎月の積立額と積立月数の合計以上で入力してください")
	}

	if input.InitialPrincipal() == 0 && input.MonthlyContribution == 0 {
		return errors.New("初期元本または毎月の積立額を入力してください")
	}

	return nil
}

func (input CalculationInput) CurrentAmount() int64 {
	return input.PrincipalAmount + input.CurrentProfit
}

func (input CalculationInput) InitialPrincipal() int64 {
	return input.PrincipalAmount - input.MonthlyContribution*int64(input.AccumulatedMonths)
}
