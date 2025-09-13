package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/repository/memory"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/controller/temperature"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Temperature/internal/handler/http"
)

func main(){
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting metadata service on port %d", port)

	r := memory.New()
	c := temperature.New(r)
	h := httpHandler.New(c)

	http.Handle("/getTemp", http.HandlerFunc(h.GetTemperature))
	http.Handle("/putTemp", http.HandlerFunc(h.PutTemperature))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
