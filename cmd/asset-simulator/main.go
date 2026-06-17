package main

import (
	"fmt"
	"os"

	"github.com/IamSBStakumi/asset-simulator/internal/app"
)

func main() {
	if err := app.RunCLI(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
