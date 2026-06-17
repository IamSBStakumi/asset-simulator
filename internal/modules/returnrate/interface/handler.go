package returnrate

import "github.com/IamSBStakumi/asset-simulator/internal/modules/returnrate/internal/domain"

type Handler struct {
	service *domain.CalculationService
}

func NewHandler() *Handler {
	calculator := domain.NewCalculator()
	service := domain.NewCalculationService(*calculator)

	return &Handler{service: service}
}

func (h *Handler) Calculate(request Request) (Response, error) {
	input := request.ToInput()

	result, err := h.service.Calculate(input)
	if err != nil {
		return Response{}, err
	}

	return NewResponse(result), nil
}
