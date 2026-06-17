package app

import (
	"fmt"
	"io"

	returnrate "github.com/IamSBStakumi/asset-simulator/internal/modules/returnrate/interface"
	simulation "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/interface"
)

func RunCLI(args []string, writer io.Writer) error {
	config, err := parseCLIArgs(args)
	if err != nil {
		return err
	}

	switch config.Mode {
	case cliModeSimulation:
		return runSimulation(config, writer)
	case cliModeReturnRate:
		return runReturnRateCalculation(config, writer)
	default:
		return fmt.Errorf("未対応のモードです: %s", config.Mode)
	}
}

func runSimulation(config cliConfig, writer io.Writer) error {
	simulator := simulation.NewHandler()

	if _, err := fmt.Fprintln(writer, "Future Assets Simulator"); err != nil {
		return err
	}

	response, err := simulator.Simulate(simulation.Request{
		Principal:           config.Principal,
		CurrentProfit:       config.CurrentProfit,
		InvestedYears:       config.InvestedYears,
		MonthlyContribution: config.MonthlyContribution,
		AnnualYieldRate:     config.AnnualYieldRate,
	})

	if err != nil {
		return err
	}

	for _, result := range response.Results {
		if _, err := fmt.Fprintf(
			writer,
			"%d年後: 総資産額 %d円、 元本 %d円、 利益 %d円\n",
			result.Year,
			result.TotalAmount,
			result.PrincipalAmount,
			result.ProfitAmount,
		); err != nil {
			return err
		}
	}

	return nil
}

func runReturnRateCalculation(config cliConfig, writer io.Writer) error {
	calculator := returnrate.NewHandler()

	if _, err := fmt.Fprintln(writer, "Return Rate Calculator"); err != nil {
		return err
	}

	response, err := calculator.Calculate(returnrate.Request{
		PrincipalAmount:     config.Principal,
		CurrentProfit:       config.CurrentProfit,
		AccumulatedMonths:   config.AccumulatedMonths,
		MonthlyContribution: config.MonthlyContribution,
	})
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(
		writer,
		"月利: %.6f%%\n年利: %.6f%%\n",
		response.MonthlyYieldRate,
		response.AnnualYieldRate,
	)

	return err
}
