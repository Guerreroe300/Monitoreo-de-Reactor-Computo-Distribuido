package http

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) PutShutdownCommand(ctx context.Context) error {
	addrs, err := g.registry.ServiceAddress(ctx, "cmd")
	if err != nil {
		return err
	}
	url := "http://" + addrs[rand.Intn(len(addrs))] + "/putCmd?cmd=Shutdown"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error creating request to DB service: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("DB service returned %d", resp.StatusCode)
		return errors.New("bad http")
	}

	return nil
}
