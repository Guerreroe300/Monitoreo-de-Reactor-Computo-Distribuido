package memory

// As we move on we will make this into a proper database
import (
	"context"
	"sync"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/repository"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
)

type Commands struct {
	sync.RWMutex

	data []*model.Temperature
}

func New() *Commands {
	return &Commands{}
}

func (r *Commands) GetLatest(_ context.Context) (*model.Temperature, error) {
	r.RLock()
	defer r.RUnlock()

	//Popping the last element
	if len(r.data) > 0 {
		m := r.data[len(r.data)-1]
		return m, nil
	} else {
		return nil, repository.ErrListEmpty
	}
}

func (r *Commands) Put(_ context.Context, temp *model.Temperature) error {
	r.Lock()
	defer r.Unlock()

	r.data = append(r.data, temp)

	return nil
}

func (r *Commands) GetAll(_ context.Context) ([]*model.Temperature, error){
	r.RLock()
	defer r.RUnlock()

	if len(r.data) > 0 {
		return r.data, nil
	} else {
		return nil, repository.ErrListEmpty
	}
}
