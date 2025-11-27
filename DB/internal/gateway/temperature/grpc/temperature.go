package grpc

import (
	"context"
	"fmt"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
)

type Gateway struct {
	registry int
}

func New(registry int) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) GetTempFromTempService(ctx context.Context) (*model.Temperature, error) {
	conn, err := grpc.NewClient("temp-service.reactor-space:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
