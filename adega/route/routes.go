package route

import (
	"github.com/gorilla/mux"
	"github.com/guiNasc/go-adega/adega/resource"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/wine/{id}", resource.Get).Methods("GET")
	router.HandleFunc("/wine", resource.GetAll).Methods("GET")
	router.HandleFunc("/wine", resource.Create).Methods("POST")
	router.HandleFunc("/wine", resource.Update).Methods("PUT")
	router.HandleFunc("/wine/{id}", resource.Delete).Methods("DELETE")
	router.HandleFunc("/wine/export", resource.Export).Methods("POST")
	return router
}

