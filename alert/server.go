package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ovas33710/wm/alert/internal/handlers"
)

func Run() {

	alertHandler, err := handlers.NewAlertHandler()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Post("/alerts", alertHandler.WriteAlert)
	r.Get("/alerts", alertHandler.ReadAlerts)

	// Server start
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: r,
	}
	log.Println("Server started...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(fmt.Sprintf("%+v", err))
	}

}
