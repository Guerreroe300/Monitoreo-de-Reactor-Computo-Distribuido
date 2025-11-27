package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/controller/commands"
	grpcHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/handler/grpc"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/handler/http"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/repository/memory"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
	"google.golang.org/grpc"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/discovery/consul"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
)

const serviceName = "cmd"

func main() {
	// name of service for docker, if running local just add this var: SERVICE_HOST=127.0.0.1
	host := os.Getenv("SERVICE_HOST")

	var port int
	var HTTPMCPORT int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.IntVar(&HTTPMCPORT, "HTTPMCPORT", 8091, "Port for MC server to getCommand")
	flag.Parse()
	log.Printf("Starting metadata service on port %d", port)
	log.Printf("Starting MC listener on port %d", HTTPMCPORT)

	// consul for docker work, before localhost
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

	// now 0.0.0.0 so it works on docker, before localhost
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

	r := memory.New()
	c := commands.New(r)
	h := grpcHandler.New(c)
	h2 := httpHandler.New(c)

	go func() {
		lis, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
		if err != nil {
			log.Fatalf("Failed to listen : %v", err)
		}
		srv := grpc.NewServer()
		gen.RegisterCommandServiceServer(srv, h)
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()

	http.Handle("/getCmd", http.HandlerFunc(h2.GetNextCommand))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", HTTPMCPORT), nil); err != nil {
		panic(err)
	}
}
