package app

import (
	"fmt"
	"io"

	simulation "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/interface"
)

func RunCLI(args []string, writer io.Writer) error {
	config, err := parseCLIArgs(args)
	if err != nil {
		return err
	}

	simulator := simulation.NewHandler()

	if _, err := fmt.Println(writer, "Future Assets Simulator"); err != nil {
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
