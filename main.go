package main

import (
	"fmt"
	"net/http"

	"github.com/domdoa/goapi/internal/handlers"
	"github.com/domdoa/goapi/internal/tools"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	database, err := tools.NewDatabase()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer func() {
		if err := database.CloseDatabase(); err != nil {
			log.Fatal("Error closing database:", err)
		}
	}()

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r, database)

	fmt.Println("Starting GO API service on port 8000...")
	err = http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
