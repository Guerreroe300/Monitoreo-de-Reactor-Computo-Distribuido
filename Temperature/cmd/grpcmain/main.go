package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/controller/temperature"
	grpc_handler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/handler/grpc"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/handler/http"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/repository/memory"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/src/gen"
	"google.golang.org/grpc"
)

//const serviceName = "temperature"

func main() {
	//host := os.Getenv("SERVICE_HOST")

	// Our port
	var port int
	var HTTPMCPORT int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.IntVar(&HTTPMCPORT, "HTTPMCPORT", 8090, "Port for MC server to upload Temp")
	flag.Parse()
	log.Printf("Starting metadata service on port %d", port)
	log.Printf("Starting MC listener on port %d", HTTPMCPORT)

	// Registry Stuff:
	//var registry *consul.Registry
	//var err error
	//if host == "localhost" {
	//	registry, err = consul.NewRegistry("localhost:8500")
	//} else {
	//	registry, err = consul.NewRegistry("dev-consul:8500")
	//}
	//if err != nil {
	//	panic(err)
	//}

	//ctx := context.Background()

	//instanceID := discovery.GenerateInstanceID(serviceName)

	//if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("%s:%d", host, port)); err != nil {
	//	panic(err)
	//}

	//go func() {
	//	for {
	//		if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
	//			log.Println("Failed to report healthy state: " + err.Error())
	//		}
	//		time.Sleep(1 * time.Second)
	//	}
	//}()

	//defer registry.Deregister(ctx, instanceID, serviceName)

	r := memory.New()
	c := temperature.New(r)
	h := grpc_handler.New(c)
	h2 := httpHandler.New(c)

	go func() {
		lis, err := net.Listen("tcp", "0.0.0.0"+":"+strconv.Itoa(port))
		if err != nil {
			log.Fatalf("Failed to listen : %v", err)
		}
		srv := grpc.NewServer()
		gen.RegisterTemperatureServiceServer(srv, h)
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()

	http.Handle("/putTemp", http.HandlerFunc(h2.PutTemperature))

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", HTTPMCPORT), nil); err != nil {
		panic(err)
	}
}
