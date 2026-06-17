package domain

type CalculationService struct {
	calc Calculator
}

func NewCalculationService(calc Calculator) *CalculationService {
	return &CalculationService{
		calc: calc,
	}
}

func (s *CalculationService) Calculate(input CalculationInput) (CalculationResult, error) {
	if err := input.Validate(); err != nil {
		return CalculationResult{}, err
	}

	return s.calc.Calculate(input)
}
