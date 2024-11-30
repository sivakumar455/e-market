package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

var log *slog.Logger

func init() {
	log = slog.New(slog.Default().Handler())
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Info("Request received for home")
	fmt.Fprintf(w, "%s", "Inventory Service")
}

func main() {

	log.Info("Inventory Main")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	err := http.ListenAndServe(":8088", mux)

	if err != nil {
		log.Error("Error: ", "error", err)
		return
	}
}
