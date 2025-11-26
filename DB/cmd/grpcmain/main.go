package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/controller/db"
	grpcHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/handler/grpc"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/repository/memory"

	temperatureGateway "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/gateway/temperature/grpc"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/discovery/consul"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
	"google.golang.org/grpc"
)

// Go subroutiner for checking in with the temp service
func checkOnTemp(c *db.Controller, ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.PutNewestTemp(ctx)
	}
}

const serviceName = "db"

func main() {
	host := os.Getenv("SERVICE_HOST")

	// OUR PORT
	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting database service on port %d", port)

	// Registry Stuff:
	var registry *consul.Registry
	var err error
	if host == "localhost" {
		registry, err = consul.NewRegistry("localhost:8500")
	} else {
		registry, err = consul.NewRegistry("dev-consul:8500")
	}
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("%s:%d", host, port)); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	tempGateway := temperatureGateway.New(registry)
	r := memory.New()
	c := db.New(r, tempGateway)
	h := grpcHandler.New(c)

	go checkOnTemp(c, ctx)

	lis, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterTemperatureServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
