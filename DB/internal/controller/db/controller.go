package db

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/repository"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/pkg/model"
)

// This in case we implement a repo somehow else later, as long as its got these functions, jala
type dbRepository interface {
	GetLatest(ctx context.Context) (*model.Temperature, error)
	Put(ctx context.Context, temp *model.Temperature) error
}

type Controller struct {
	repo dbRepository
}

func New(repo dbRepository) *Controller {
	return &Controller{repo: repo}
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

func (c *Controller) PutNewestTemp() error {
	url := "http://localhost:8081/getTemp"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error creating request to Temp service: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Temp service returned %d", resp.StatusCode)
		return repository.ErrHttpIssue
	}

	var temp model.Temperature
	if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return err
	}

	c.Put(resp.Request.Context(), &temp)
	return nil
}
