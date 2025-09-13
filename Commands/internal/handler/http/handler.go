package http

import (
	"errors"
	"log"
	"net/http"
	"encoding/json"

	commands "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/controller/commands"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/repository"
)

type Handler struct{
	ctrl *commands.Controller
}

func New(ctrl *commands.Controller) *Handler{
	return &Handler{ctrl: ctrl}
}

func (h* Handler) GetNextCommand(w http.ResponseWriter, req *http.Request) {
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

func (h* Handler) PutNewCommand(w http.ResponseWriter, req *http.Request) {
	cmd:= req.FormValue("cmd")
	if(cmd == ""){
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx := req.Context()

	err := h.ctrl.Put(ctx, &cmd)

	if (err != nil){
		log.Printf("Error Putting: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}else {
		w.WriteHeader(http.StatusCreated)
	}

}
