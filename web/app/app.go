package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	router     *mux.Router
	controller *Controller
	service    *Service
	db         *DB
	server     *http.Server
}

func NewApp() App {
	db := NewDB()
	service := Service{db}

	controller := &Controller{service}

	router := InitializeRouter(*controller)

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	return App{router, controller, &service, &db, server}
}

func (app App) Start(port string, finished chan bool) error {
	err := app.server.ListenAndServe()
	if err != nil {
		return err
	}
	finished <- true
	return nil
}
