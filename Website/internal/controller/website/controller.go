package website

import (
	"context"

	model "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
)

type dbGateway interface {
	GetAllFromDBService(ctx context.Context) ([]*model.Temperature, error)
}

type cmdGateway interface {
	PutShutdownCommand(ctx context.Context) error
}

type Controller struct {
	dbGate  dbGateway
	cmdGate cmdGateway
}

func New(dbGate dbGateway, cmdGate cmdGateway) *Controller {
	return &Controller{dbGate: dbGate, cmdGate: cmdGate}
}

func (c *Controller) GetAllDB(ctx context.Context) ([]*model.Temperature, error) {
	temp, err := c.dbGate.GetAllFromDBService(ctx)

	if err != nil {
		return nil, err
	}

	return temp, nil
}

func (c *Controller) ShutdownButton(ctx context.Context) error {
	err := c.cmdGate.PutShutdownCommand(ctx)

	if err != nil {
		return err
	}

	return nil
}
