package http

import (
	"errors"
	"log"
	"net/http"
	"encoding/json"
	"strconv"

	temperature "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/controller/temperature"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/repository"
)

type Handler struct{
	ctrl *temperature.Controller
}

func New(ctrl *temperature.Controller) *Handler{
	return &Handler{ctrl: ctrl}
}

func (h* Handler) GetTemperature(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	m, err := h.ctrl.Get(ctx)

	if (err != nil && errors.Is(err, repository.ErrNotFound)){
		w.WriteHeader(http.StatusNotFound)
		return
	} else if (err != nil){
		log.Printf("Repository error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil{
		log.Printf("Response Error: v%\n", err)
	}

}

func (h* Handler) PutTemperature(w http.ResponseWriter, req *http.Request) {
	temp:= req.FormValue("temp")
	tempf, conErr := strconv.ParseFloat(temp, 32)
	temp32 := float32(tempf)
	if(temp == "" || conErr != nil){
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx := req.Context()

	err := h.ctrl.Put(ctx, temp32)

	if (err != nil){
		log.Printf("Error Putting: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}else {
		w.WriteHeader(http.StatusCreated)
	}

}
