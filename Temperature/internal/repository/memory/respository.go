package memory

import (
	"context"
	"sync"
	"time"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/repository"
	model "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
)

type Repository struct{
	sync.RWMutex
	data *model.Temperature
}

func New() *Repository{
	return &Repository{data: &model.Temperature{Temperature:10.0, Date: time.Now()}}
}

func (r *Repository) Get(_ context.Context) (*model.Temperature, error){
	r.RLock()
	defer r.RUnlock()

	m := r.data

	if m == nil{
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) Put(_ context.Context, temp float32) (error){
	r.Lock()
	defer r.Unlock()

	if (r.data == nil){
		r.data = &model.Temperature{}
	}

	r.data.Temperature = temp
	r.data.Date = time.Now()

	return nil
}
