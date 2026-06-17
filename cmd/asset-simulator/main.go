package main

import (
	"fmt"
	"os"

	"github.com/IamSBStakumi/asset-simulator/internal/app"
)

func main() {
	if err := app.RunCLI(os.Args[1:], os.Stdout); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}
}
