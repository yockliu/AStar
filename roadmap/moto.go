package roadmap

import (
	"fmt"
)

var KeyMotoCounter = 1

type Moto struct {
	ID    string `json:"id,omitempty"`
	Point `json:"point,omitempty"`
}

func NewMoto(x float64, y float64) *Moto {
	m := new(Moto)
	m.ID = fmt.Sprintf("%6d", KeyMotoCounter)
	m.Point.X = x
	m.Point.Y = y
	KeyMotoCounter = KeyMotoCounter + 1
	return m
}
