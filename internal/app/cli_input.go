package app

import "flag"

// cliConfig はCLI引数から組み立てた設定値です。
type cliConfig struct {
	Principal           int64
	CurrentProfit       int64
	InvestedYears       int
	MonthlyContribution int64
	AnnualYieldRate     float64
	TargetYears         []int
}

// parseCLIArgs はCLI引数を解析します。
func parseCLIArgs(args []string) (cliConfig, error) {
	flagSet := flag.NewFlagSet("asset-simulator", flag.ContinueOnError)

	principal := flagSet.Int64("principal", 0, "元本")
	currentProfit := flagSet.Int64("current-profit", 0, "現在の利益。含み損がある場合は負の値を指定できます")
	investedYears := flagSet.Int("invested-years", 0, "これまでの積立年数")
	monthlyContribution := flagSet.Int64("monthly-contribution", 0, "毎月の積立額")
	annualYieldRate := flagSet.Float64("annual-yield-rate", 0, "年利。5%の場合は5を指定します")

	if err := flagSet.Parse(args); err != nil {
		return cliConfig{}, err
	}

	return cliConfig{
		Principal:           *principal,
		CurrentProfit:       *currentProfit,
		InvestedYears:       *investedYears,
		MonthlyContribution: *monthlyContribution,
		AnnualYieldRate:     *annualYieldRate,

		// 出力対象年はdomain側のデフォルト値を使うため、ここでは未指定にします。
		TargetYears: nil,
	}, nil
}
