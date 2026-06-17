package domain

type Simulator struct{}

func NewSimulator() *Simulator {
	return &Simulator{}
}

func (s *Simulator) Simulate(input SimulationInput) (SimulationResult, error) {
	// validation
	normalizedInput := input.Normalize()
	if err := normalizedInput.Validate(); err != nil {
		return SimulationResult{}, err
	}

	// calculate

	return SimulationResult{}, nil
}

func max(values []int) int {
	maxValue := values[0]

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}