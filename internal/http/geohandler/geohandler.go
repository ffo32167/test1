package geohandler

import (
	"encoding/json"
	"net/http"

	"github.com/ffo32167/task1/internal"
	"go.uber.org/zap"
)

type GeoHandler struct {
	log *zap.Logger
}

func New(log *zap.Logger) GeoHandler {
	return GeoHandler{log: log}
}

func (h GeoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var uPos internal.UserGeoPosition
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&uPos)
	if err != nil {
		h.log.Error("geohandler: cant decode user position:", zap.Error(err))
	}

	data, err := internal.Process(req.Context(), uPos)
	if err != nil {
		h.log.Error("geohandler: cant process user position:", zap.Error(err))
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		h.log.Error("geohandler: cant encode response:", zap.Error(err))
	}
}
