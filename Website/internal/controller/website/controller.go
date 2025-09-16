package website

import (
	"context"

	model "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Get5Latest(ctx context.Context) ([]*model.Temperature, error) {
	// Here we use the 5 latest from the DB service
	return []*model.Temperature{}, nil
}
