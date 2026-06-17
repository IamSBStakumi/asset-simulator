package domain

import (
	"errors"
	"math"
)

const (
	lowerMonthlyRate = -0.999999999999
	initialUpperRate = 0.01
	maxMonthlyRate   = 1000.0
	searchIterations = 200
	equalTolerance   = 0.000001
)

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (*Calculator) Calculate(input CalculationInput) (CalculationResult, error) {
	targetAmount := float64(input.CurrentAmount())
	zeroRateAmount := calculateAmount(input, 0)

	if nearlyEqual(targetAmount, zeroRateAmount) {
		return buildResult(0), nil
	}

	lowerAmount := calculateAmount(input, lowerMonthlyRate)
	if targetAmount < lowerAmount {
		return CalculationResult{}, errors.New("現在資産額が小さすぎるため利回りを計算できません")
	}

	upperRate, err := findUpperRate(input, targetAmount)
	if err != nil {
		return CalculationResult{}, err
	}

	monthlyRate := searchMonthlyRate(input, targetAmount, lowerMonthlyRate, upperRate)

	return buildResult(monthlyRate), nil
}

func calculateAmount(input CalculationInput, monthlyRate float64) float64 {
	amount := float64(input.InitialPrincipal())

	for month := 1; month <= input.AccumulatedMonths; month++ {
		amount = amount*(1+monthlyRate) + float64(input.MonthlyContribution)
	}

	return amount
}

func findUpperRate(input CalculationInput, targetAmount float64) (float64, error) {
	upperRate := initialUpperRate

	for calculateAmount(input, upperRate) < targetAmount {
		upperRate *= 2
		if upperRate > maxMonthlyRate {
			return 0, errors.New("利回りが大きすぎるため計算できません")
		}
	}

	return upperRate, nil
}

func searchMonthlyRate(
	input CalculationInput,
	targetAmount float64,
	lowerRate float64,
	upperRate float64,
) float64 {
	for range searchIterations {
		middleRate := (lowerRate + upperRate) / 2
		amount := calculateAmount(input, middleRate)

		if amount < targetAmount {
			lowerRate = middleRate
			continue
		}

		upperRate = middleRate
	}

	return (lowerRate + upperRate) / 2
}

func buildResult(monthlyRate float64) CalculationResult {
	return CalculationResult{
		MonthlyYieldPercentage: monthlyRate * 100,
		AnnualYieldPercentage:  (math.Pow(1+monthlyRate, 12) - 1) * 100,
	}
}

func nearlyEqual(left float64, right float64) bool {
	return math.Abs(left-right) < equalTolerance
}
