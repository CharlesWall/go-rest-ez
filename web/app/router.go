package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter(controller Controller) *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.HandlerFunc(controller.handleRequest))

	return router
}
