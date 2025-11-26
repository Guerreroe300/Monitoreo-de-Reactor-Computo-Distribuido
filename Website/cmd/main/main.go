package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Website/internal/controller/website"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Website/internal/handler/http"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/discovery/consul"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"

	cmdGateway "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Website/internal/gateway/commands/http"
	dbGateway "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Website/internal/gateway/db/http"
)

const serviceName = "website"

func main() {
	host := os.Getenv("SERVICE_HOST")

	var port int
	flag.IntVar(&port, "port", 8080, "API handler port")
	flag.Parse()
	log.Printf("Starting website service on port %d", port)

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

	dbGate := dbGateway.New(registry)
	cmdGate := cmdGateway.New(registry)

	c := website.New(dbGate, cmdGate)
	h := httpHandler.New(c)

	// Serve the HTML
	http.Handle("/", http.HandlerFunc(h.MainHtml))

	// API endpoint: return table rows
	http.Handle("/api/data", http.HandlerFunc(h.TableGet))

	// API endpoint: button action
	http.HandleFunc("/api/doSomething", http.HandlerFunc(h.ButtonHandler))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
