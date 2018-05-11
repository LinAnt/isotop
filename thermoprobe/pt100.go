package thermoprobe

import "math/rand"

type PT100 struct {
	Thermoprobe
}

func NewPT100() *PT100 {
	p := &PT100{}
	return p
}

func Read() (float32, error) {
	return rand.Float32(), nil
}
