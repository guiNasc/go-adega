/*package resource

import (
	"net/http"
	"encoding/json"
	"github.com/guiNasc/go-adega/adega/model"
	"github.com/gorilla/mux"
	"github.com/guiNasc/go-adega/adega/service"
)

func setCotentType(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return w
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w = setCotentType(w)
	json.NewEncoder(w).Encode(service.GetAll_)
}*/