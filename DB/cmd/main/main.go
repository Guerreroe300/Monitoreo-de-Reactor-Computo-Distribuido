package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/controller/db"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/handler/http"
	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/DB/internal/repository/memory"
)

// Go subroutiner for checking in with the temp service
func checkOnTemp(c *db.Controller) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.PutNewestTemp()
	}
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting database service on port %d", port)

	r := memory.New()
	c := db.New(r)
	h := httpHandler.New(c)

	go checkOnTemp(c)

	http.Handle("/dbGetLatest", http.HandlerFunc(h.GetLatestTemp))
	http.Handle("/getAll", http.HandlerFunc(h.GetAllTemps))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
