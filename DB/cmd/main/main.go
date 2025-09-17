package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/controller/db"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/handler/http"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/repository/memory"

	temperatureGateway "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/gateway/temperature/http"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/discovery/consul"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
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

	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting database service on port %d", port)

	// Registry Stuff:
	registry, err := consul.NewRegistry("dev-consul:8500")
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
	h := httpHandler.New(c)

	go checkOnTemp(c, ctx)

	http.Handle("/dbGetLatest", http.HandlerFunc(h.GetLatestTemp))
	http.Handle("/getAll", http.HandlerFunc(h.GetAllTemps))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
