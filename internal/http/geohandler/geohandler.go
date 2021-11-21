package geohandler

import (
	"encoding/json"
	"net/http"

	"github.com/ffo32167/test1/internal"
	"go.uber.org/zap"
)

type GeoHandler struct {
	storage internal.Storage
	log     *zap.Logger
}

func New(log *zap.Logger, storage internal.Storage) GeoHandler {
	return GeoHandler{log: log, storage: storage}
}

func (h GeoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var uPos internal.UserGeoPosition
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&uPos)
	if err != nil {
		h.log.Error("geohandler: cant decode user position:", zap.Error(err))
	}

	_, len, err := internal.Process(h.storage, internal.UserGeoPosition{})
	if err != nil {
		h.log.Error("geohandler: cant encode response:", zap.Error(err))
	}

	h.log.Info("Slice, ", zap.Int("len", len))
	/*	err = json.NewEncoder(res).Encode(data)
		if err != nil {
			h.log.Error("geohandler: cant encode response:", zap.Error(err))
		}
	*/
}
