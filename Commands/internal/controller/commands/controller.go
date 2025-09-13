package commands 

import (
	"context"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/repository"
)

// This in case we implement a repo somehow else later, as long as its got these functions, jala
type commandRepository interface{
	Pop(ctx context.Context) (*string, error)	
	Append(ctx context.Context, comand *string) (error)
}

type Controller struct{
	repo commandRepository
}

func New(repo commandRepository) *Controller{
	return &Controller{repo: repo}
}

func (c* Controller) Get(ctx context.Context) (*string, error){
	res, error := c.repo.Pop(ctx)

	if (error != nil){
		return nil, repository.ErrNotFound
	}

	return res, error
}

func (c* Controller) Put(ctx context.Context, comand *string) (error){
	error := c.repo.Append(ctx, comand)

	return error 
}
