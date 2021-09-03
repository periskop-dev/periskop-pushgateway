package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	//"github.com/soundcloud/periskop-go"
)

func main() {
	var (
		port = flag.String("port", os.Getenv("PORT"), "The server port")
	)

	flag.Parse()

	router := mux.NewRouter()
	repo := sync.Map{}

	// API routing
	setupAPIRouting(&repo, router)

	// Telemetry endpoints
	//errorExporter := periskop.NewErrorExporter(&metrics.ErrorCollector)
	//periskopHandler := periskop.NewHandler(errorExporter)

	//http.Handle("/errors", periskopHandler)
	http.HandleFunc("/-/health", healthHandler)

	address := fmt.Sprintf(":%s", *port)
	log.Printf("Serving on address %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func setupAPIRouting(repo *sync.Map, r *mux.Router) {
	r.Handle("/errors/{target_name}/", NewErrorsGatewayHandler(repo)).Methods(http.MethodPost)
	r.Handle("/errors/", NewErrorsListHandler(repo)).Methods(http.MethodGet)
	//r.Use(api.CORSLocalhostMiddleware(r))
	http.Handle("/", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatalf("error running health handler")
	}
}
