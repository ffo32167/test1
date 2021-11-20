package http

import (
	"net/http"
	"time"

	"github.com/ffo32167/task1/internal/http/geohandler"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ApiServer struct {
	port    string
	timeout time.Duration
	log     *zap.Logger
}

func New(port string, timeout time.Duration, log *zap.Logger) ApiServer {
	return ApiServer{port: port, timeout: timeout, log: log}
}

func (as ApiServer) Run() error {
	geoHandler := geohandler.New(as.log)

	router := mux.NewRouter()
	router.Handle("/", geoHandler).Methods("GET")

	err := http.ListenAndServe(as.port, router)
	if err != nil {
		return err
	}
	return nil
}
