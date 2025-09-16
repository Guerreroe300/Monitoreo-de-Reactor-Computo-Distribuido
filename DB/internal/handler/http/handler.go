package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/controller/db"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/repository"
)

type Handler struct {
	ctrl *db.Controller
}

func New(ctrl *db.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetLatestTemp(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	m, err := h.ctrl.GetLatest(ctx)

	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response Error: v%\n", err)
	}

}

func (h* Handler) GetAllTemps(w http.ResponseWriter, req *http.Request){
	ctx := req.Context()

	m, err := h.ctrl.GetAll(ctx)

	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response Error: v%\n", err)
	}
}

// nvm this function, i have to manually request from the interface every once in a while
/*func (h *Handler) PutNewCommand(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	tem
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx := req.Context()

	err := h.ctrl.Put(ctx, id)

	if err != nil {
		log.Printf("Error Putting: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}*/
