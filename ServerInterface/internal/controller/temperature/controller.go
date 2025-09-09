package temperature

import (
	"context"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/ServerInterface/internal/repository"
	model "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/ServerInterface/pkg/model"
)

// This in case we implement a repo somehow else later, as long as its got these functions, jala
type temperatureRepository interface{
	Get(ctx context.Context) (*model.Temperature, error)	
	Put(ctx context.Context, temp float32) (error)
}

type Controller struct{
	repo temperatureRepository
}

func New(repo temperatureRepository) *Controller{
	return &Controller{repo: repo}
}

func (c* Controller) Get(ctx context.Context) (*model.Temperature, error){
	res, error := c.repo.Get(ctx)

	if (error != nil){
		return nil, repository.ErrNotFound
	}

	return res, error
}

func (c* Controller) Put(ctx context.Context, temp float32) (error){
	error := c.repo.Put(ctx, temp)

	return error 
}
