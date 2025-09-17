package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
	discovery "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/pkg/registry"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) GetAllFromDBService(ctx context.Context) ([]*model.Temperature, error) {
	addrs, err := g.registry.ServiceAddress(ctx, "db")
	if err != nil {
		return nil, err
	}
	url := "http://" + addrs[rand.Intn(len(addrs))] + "/getAll"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error creating request to DB service: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("DB service returned %d", resp.StatusCode)
		return nil, errors.New("bad http")
	}

	var temp []*model.Temperature
	if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return nil, err
	}

	return temp, nil
}
