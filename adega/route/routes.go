package route

import (
	"github.com/gorilla/mux"
	"github.com/guiNasc/go-adega/adega/service"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/wine/{id}", service.Get).Methods("GET")
	router.HandleFunc("/wine", service.GetAll).Methods("GET")
	router.HandleFunc("/wine", service.Create).Methods("POST")
	router.HandleFunc("/wine", service.Update).Methods("PUT")
	router.HandleFunc("/wine/{id}", service.Delete).Methods("DELETE")
	router.HandleFunc("/wine/export", service.Export).Methods("POST")
	return router
}

