package internal

import (
	"sync"
)

type UserGeoPosition struct {
	Id    [16]byte // https://github.com/google/uuid
	Coord Coord
}

type Coord struct {
	Lat  float64
	Long float64
}

type Locations struct {
	sync.RWMutex
	uPos []UserGeoPosition
}

type Storage interface {
	Save(newCoords UserGeoPosition) (int, error)
	FindLocation(newCoords UserGeoPosition) (Coord, error)
}

func Process(storage Storage, newCoords UserGeoPosition) (Coord, int, error) {
	loc, err := storage.FindLocation(newCoords)
	if err != nil {
		return Coord{}, 0, err
	}
	count, err := storage.Save(newCoords)
	if err != nil {
		return Coord{}, 0, err
	}
	return loc, count, nil
}
