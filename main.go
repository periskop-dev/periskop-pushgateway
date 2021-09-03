package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/soundcloud/periskop-go"
)

func main() {
	var (
		port = flag.String("port", os.Getenv("PORT"), "The server port")
	)

	flag.Parse()

	router := mux.NewRouter()

	// API routing
	collector := periskop.NewErrorCollector()
	setupAPIRouting(&collector, router)

	// Telemetry endpoints
	errorExporter := periskop.NewErrorExporter(&collector)
	periskopHandler := periskop.NewHandler(errorExporter)

	http.Handle("/-/errors", periskopHandler)
	http.HandleFunc("/-/health", healthHandler)

	address := fmt.Sprintf(":%s", *port)
	log.Printf("Serving on address %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func setupAPIRouting(collector *periskop.ErrorCollector, r *mux.Router) {
	r.Handle("/errors", NewErrorsGatewayHandler(collector)).Methods(http.MethodPost)
	http.Handle("/", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatalf("error running health handler")
	}
}
