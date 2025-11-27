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

func (g *Gateway) GetAllFromDBService(ctx context.Context) ([]*model.Temperature, error) {
	conn, err := grpc.NewClient("db-service.reactor-space:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewTemperatureServiceClient(conn)

	resp, err := client.GetAllTemperatures(ctx, &emptypb.Empty{})

	if err != nil {
		fmt.Printf("Error creating request to DB service: %v\n", err)
		return nil, err
	}

	protoTemps := resp.GetTemperatureReadings()
	goTemps := make([]*model.Temperature, len(protoTemps))
	for i, protoTemp := range protoTemps {
		// Assuming ProtoToTemperature converts *gen.Temperature to *model.Temperature
		goTemps[i] = model.TemperatureFromProto(protoTemp)
	}

	return goTemps, nil
}
