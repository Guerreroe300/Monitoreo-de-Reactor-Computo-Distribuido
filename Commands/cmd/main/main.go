package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/repository/memory"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/controller/commands"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Commands/internal/handler/http"
)

func main(){
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()
	log.Printf("Starting metadata service on port %d", port)

	r := memory.New()
	c := commands.New(r)
	h := httpHandler.New(c)

	http.Handle("/getCmd", http.HandlerFunc(h.GetNextCommand))
	http.Handle("/putCmd", http.HandlerFunc(h.PutNewCommand))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
