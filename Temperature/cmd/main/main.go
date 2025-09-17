package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/controller/temperature"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/handler/http"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/repository/memory"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/discovery/consul"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
)

const serviceName = "temperature"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting metadata service on port %d", port)

	// Registry Stuff:
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
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
	c := temperature.New(r)
	h := httpHandler.New(c)

	http.Handle("/getTemp", http.HandlerFunc(h.GetTemperature))
	http.Handle("/putTemp", http.HandlerFunc(h.PutTemperature))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
