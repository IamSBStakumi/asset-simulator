package simulation

import "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/internal/domain"

type Handler struct {
	service *domain.Simulator
}

func NewHandler() *Handler {
	calculator := domain.NewCalculator()
	service := domain.NewSimulator(*calculator)

	return &Handler{service: service}
}

func (h *Handler) Simulate(request Request) (Response, error) {
	input := request.ToInput()

	result, err := h.service.Simulate(input)
	if err != nil {
		return Response{}, err
	}

	return NewResponse(result), nil
}
