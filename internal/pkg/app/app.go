package app

import (
	"errors"
	"net/http"
	"os"
	"quest/internal/controller"
	"quest/internal/pkg/logger"
	"quest/internal/service"
)

type App struct {
	os  service.OrderService
	rs  service.RoomAvailabilityService
	mux *http.ServeMux
}

func New() (*App, error) {
	a := &App{}
	a.rs = service.NewRoom()
	a.os = service.NewOrder(a.rs)
	handler := controller.NewHTTPHandler(a.os, a.rs)
	a.mux = http.NewServeMux()
	a.mux.HandleFunc("/orders", handler.CreateOrder)
	return a, nil
}

func (a *App) Run() error {
	logger.LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", a.mux)
	if errors.Is(err, http.ErrServerClosed) {
		logger.LogErrorf("Server closed")
	} else if err != nil {
		logger.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
	return nil
}
