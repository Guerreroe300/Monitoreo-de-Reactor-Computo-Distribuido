package website

import (
	"context"
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	
	model "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetAllDB(ctx context.Context) ([]*model.Temperature, error) {
	url := "http://localhost:8083/getAll"

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
