package grpc

import (
	"context"
	"fmt"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/grpcutil"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) GetTempFromTempService(ctx context.Context) (*model.Temperature, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "temperature", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := gen.NewTemperatureServiceClient(conn)

	resp, err := client.GetLatestTemperature(ctx, &emptypb.Empty{})

	if err != nil {
		fmt.Printf("Error creating request to Temp service: %v\n", err)
		return nil, err
	}

	return model.TemperatureFromProto(resp.TemperatureReading), nil
}
