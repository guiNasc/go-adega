package resource

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/guiNasc/go-adega/adega/model"
	"github.com/guiNasc/go-adega/adega/service"
)

func setCotentType(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return w
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w = setCotentType(w)
	json.NewEncoder(w).Encode(service.GetAll_())
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	w = setCotentType(w)
	var wines = service.Get_(id)
	json.NewEncoder(w).Encode(wines)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var wine model.Wine
	_ = json.NewDecoder(r.Body).Decode(&wine)
	service.Create_(wine)
	w = setCotentType(w)
	json.NewEncoder(w).Encode(wine)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	service.Delete_(id)
	w = setCotentType(w)
	json.NewEncoder(w).Encode(model.Wines)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var wine model.Wine
	_ = json.NewDecoder(r.Body).Decode(&wine)
	w = setCotentType(w)
	json.NewEncoder(w).Encode(service.Update_(wine))
}

func Export(w http.ResponseWriter, r *http.Request) {
	w = setCotentType(w)
	json.NewEncoder(w).Encode(service.GetAll_())
}