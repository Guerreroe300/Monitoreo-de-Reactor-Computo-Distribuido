package grpc

import (
	"context"
	"fmt"

	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/grpcutil"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) PutShutdownCommand(ctx context.Context) error {
	conn, err := grpcutil.ServiceConnection(ctx, "cmd", g.registry)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := gen.NewCommandServiceClient(conn)

	req := &gen.PutNewCommandRequest{
		Command: "Shutdown", // Populate the Command field
	}
	_, erra := client.PutNewCommand(ctx, req)

	if erra != nil {
		fmt.Printf("Error creating request to DB service: %v\n", erra)
		return erra
	}

	return nil
}
