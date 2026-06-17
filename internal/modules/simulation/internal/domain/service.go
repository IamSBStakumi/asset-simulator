package domain

type Simulator struct {
	calc Calculator
}

func NewSimulator(calc Calculator) *Simulator {
	return &Simulator{
		calc: calc,
	}
}

func (s *Simulator) Simulate(input SimulationInput) (SimulationResult, error) {
	// validation
	normalizedInput := input.Normalize()
	if err := normalizedInput.Validate(); err != nil {
		return SimulationResult{}, err
	}

	// calculate
	rowsByYear := s.calc.Calculate(normalizedInput)

	return SimulationResult{
		Rows: rowsByYear,
	}, nil
}
