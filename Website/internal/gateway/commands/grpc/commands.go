package grpc

import (
	"context"
	"fmt"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	registry int
}

func New(registry int) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) PutShutdownCommand(ctx context.Context) error {
	conn, err := grpc.NewClient("commands-service.reactor-space:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
