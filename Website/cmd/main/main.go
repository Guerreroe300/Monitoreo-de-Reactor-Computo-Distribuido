package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Website/internal/controller/website"
	httpHandler "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Website/internal/handler/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "API handler port")
	flag.Parse()
	log.Printf("Starting website service on port %d", port)

	c := website.New()
	h := httpHandler.New(c)

	// Serve the HTML
	http.Handle("/", http.HandlerFunc(h.MainHtml))

	// API endpoint: return table rows
	http.Handle("/api/data", http.HandlerFunc(h.TableGet))

	// API endpoint: button action
	http.HandleFunc("/api/doSomething", http.HandlerFunc(h.ButtonHandler))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
