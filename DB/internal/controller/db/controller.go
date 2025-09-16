package db

import (
	"context"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/repository"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
)

// This in case we implement a repo somehow else later, as long as its got these functions, jala
type dbRepository interface {
	GetLatest(ctx context.Context) (*model.Temperature, error)
	Put(ctx context.Context, temp *model.Temperature) error
	GetAll(_ context.Context) ([]*model.Temperature, error)
}

// Interface for a gateway to do external shit
type temperatureGateway interface {
	GetTempFromTempService(ctx context.Context) (*model.Temperature, error)
}

type Controller struct {
	repo        dbRepository
	tempGateway temperatureGateway
}

func New(repo dbRepository, tempGateway temperatureGateway) *Controller {
	return &Controller{repo: repo, tempGateway: tempGateway}
}

func (c *Controller) GetLatest(ctx context.Context) (*model.Temperature, error) {
	res, error := c.repo.GetLatest(ctx)

	if error != nil {
		return nil, repository.ErrNotFound
	}

	return res, error
}

func (c *Controller) Put(ctx context.Context, temp *model.Temperature) error {
	error := c.repo.Put(ctx, temp)

	return error
}

func (c *Controller) PutNewestTemp(ctx context.Context) error {
	resp, err := c.tempGateway.GetTempFromTempService(ctx)

	if err != nil {
		return err
	}

	c.Put(ctx, resp)
	return nil
}

func (c *Controller) GetAll(ctx context.Context) ([]*model.Temperature, error) {
	temps, err := c.repo.GetAll(ctx)

	if err != nil {
		return nil, repository.ErrListEmpty
	}

	return temps, nil
}
