package entities

import (
	"errors"
)

type Geo struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius uint16     `json:"radius"`
}

func (g *Geo) Validate() error {
	if g.Lat==0 {
		return errors.New("lat could not be zero")
	}
	if g.Lon==0 {
		return errors.New("lon could not be zero")
	}
	if g.Radius==0 {
		return errors.New("radius could not be zero")
	}
	return nil
}