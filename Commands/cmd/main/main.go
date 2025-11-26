package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/controller/commands"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/handler/http"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/repository/memory"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/discovery/consul"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
)

const serviceName = "cmd"

func main() {
	// name of service for docker, if running local just add this var: SERVICE_HOST=127.0.0.1
	host := os.Getenv("SERVICE_HOST")

	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()
	log.Printf("Starting metadata service on port %d", port)

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
	h := httpHandler.New(c)

	http.Handle("/getCmd", http.HandlerFunc(h.GetNextCommand))
	http.Handle("/putCmd", http.HandlerFunc(h.PutNewCommand))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
