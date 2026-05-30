package health

import "context"

type CheckRequest struct{}

type CheckResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

type CheckHandler struct {
	serviceName string
}

func NewCheckHandler(serviceName string) *CheckHandler {
	return &CheckHandler{
		serviceName: serviceName,
	}
}

func (h *CheckHandler) Handle(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return &CheckResponse{
		Service: h.serviceName,
		Status:  "ok",
	}, nil
}
