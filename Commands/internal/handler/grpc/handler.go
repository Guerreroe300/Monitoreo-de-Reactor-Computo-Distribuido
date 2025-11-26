package grpc_handler

import (
	"context"
	"errors"
	"log"

	commands "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/controller/commands"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/repository"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	gen.UnimplementedCommandServiceServer
	ctrl *commands.Controller
}

func New(ctrl *commands.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetNextCommand(ctx context.Context, req *emptypb.Empty) (*gen.GetNextCommandResponse, error) {
	m, err := h.ctrl.Get(ctx)

	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "temperature reading not found: %v", err)
	} else if err != nil {
		log.Printf("Repository error: %v\n", err)
		return nil, status.Errorf(codes.Internal, "failed to retrieve temperature reading: %v", err)
	}

	commandMessage := &gen.Command{
		// Set the string field inside the Command message
		Command: *m,
	}
	// me quede en retornal el model
	return &gen.GetNextCommandResponse{Command: commandMessage}, nil
}

func (h *Handler) PutNewCommand(ctx context.Context, req *gen.PutNewCommandRequest) (*emptypb.Empty, error) {
	cmd := req.Command

	if cmd == "" {
		return nil, status.Errorf(codes.Internal, "temperature reading not found: bad request")
	}

	err := h.ctrl.Put(ctx, &cmd)

	if err != nil {
		log.Printf("Error Putting: %v\n", err)
		return nil, status.Errorf(codes.Internal, "temperature reading not found: bad request")
	} else {
		return &emptypb.Empty{}, nil
	}
}
