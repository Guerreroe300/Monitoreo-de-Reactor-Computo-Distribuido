package grpcutil

import (
	"context"
	"math/rand"

	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, registry discovery.Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.ServiceAddress(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	return grpc.NewClient(addrs[rand.Intn(len(addrs))], grpc.WithTransportCredentials(insecure.NewCredentials()))
}
