package commandmemory 

import (
	"sync"
	"context"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/ServerInterface/internal/repository"
)

type Commands struct{
	sync.RWMutex
	data []*string 	
}

func New() *Commands {
	return &Commands{}
}

func (r *Commands) Pop(_ context.Context) (*string, error){
	r.RLock()
	defer r.RUnlock()
	
	//Popping the last element
	if len(r.data) > 0 {
		m := r.data[len(r.data)-1]
		r.data = r.data[:len(r.data)-1]
		return m, nil
	} else {
		return nil, repository.ErrListEmpty
	}
}

func (r *Commands) Append(_ context.Context) (*string, error){
	r.RLock()
	defer r.RUnlock()
	
	//Popping the last element
	if len(r.data) > 0 {
		m := r.data[len(r.data)-1]
		r.data = r.data[:len(r.data)-1]
		return m, nil
	} else {
		return nil, repository.ErrListEmpty
	}
}
