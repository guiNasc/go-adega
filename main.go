package main

import (
	"log"
	"net/http"
	"github.com/guiNasc/go-adega/adega/service"
	"github.com/guiNasc/go-adega/adega/route"
)

func main() {
	service.InitServiceMock()
	router := route.Routes()
	log.Fatal(http.ListenAndServe(":8000", router))
}

