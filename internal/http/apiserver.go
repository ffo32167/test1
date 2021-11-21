package http

import (
	"net/http"

	"github.com/ffo32167/test1/internal"
	"github.com/ffo32167/test1/internal/http/geohandler"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ApiServer struct {
	port string
	log  *zap.Logger
}

func New(port string, log *zap.Logger) ApiServer {
	return ApiServer{port: port, log: log}
}

func (as ApiServer) Run(storage internal.Storage) error {
	geoHandler := geohandler.New(as.log, storage)

	router := mux.NewRouter()
	router.Handle("/", geoHandler).Methods("GET")

	err := http.ListenAndServe(as.port, router)
	if err != nil {
		return err
	}
	return nil
}
