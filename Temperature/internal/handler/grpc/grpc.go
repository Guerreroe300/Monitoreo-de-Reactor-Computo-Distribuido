package grpc_handler

import (
	"context"
	"errors"
	"log"

	temperature "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/controller/temperature"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/repository"
	model "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedTemperatureServiceServer
	ctrl *temperature.Controller
}

func New(ctrl *temperature.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetTemperature(ctx context.Context) (*gen.GetSingleTemperatureResponse, error) {
	m, err := h.ctrl.Get(ctx)

	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "temperature reading not found: %v", err)
	} else if err != nil {
		log.Printf("Repository error: %v\n", err)
		return nil, status.Errorf(codes.Internal, "failed to retrieve temperature reading: %v", err)
	}

	return &gen.GetSingleTemperatureResponse{TemperatureReading: model.TemperatureToProto(m)}, nil
}
