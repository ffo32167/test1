package slice

import (
	"sync"

	"github.com/ffo32167/test1/internal"
)

type Slice struct {
	UserCoordinates *[]internal.UserGeoPosition
	*sync.Mutex
}

func New() Slice {
	slice := make([]internal.UserGeoPosition, 0)
	return Slice{UserCoordinates: &slice}
}

func (s Slice) Save(newCoords internal.UserGeoPosition) (int, error) {
	*s.UserCoordinates = append(*s.UserCoordinates, internal.UserGeoPosition{})
	return len(*s.UserCoordinates), nil
}

func (s Slice) FindLocation(newCoords internal.UserGeoPosition) (internal.Coord, error) {
	return internal.Coord{Lat: newCoords.Coord.Lat, Long: newCoords.Coord.Long}, nil
}
