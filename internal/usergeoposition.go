package internal

import "context"

type UserGeoPosition struct {
	Id    [16]byte // https://github.com/google/uuid
	Coord Coord
}

type Coord struct {
	Lat  float64
	Long float64
}

func Process(ctx context.Context, upos UserGeoPosition) (Coord, error) {
	return Coord{}, nil
}
