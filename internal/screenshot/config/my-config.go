package config

import (
	"ArenaStatus/internal/screenshot/plate"
	"ArenaStatus/internal/types"
)

type MyConfig struct {
	startPos    plate.Plate
	shortOffset types.Offset
	longOffset types.Offset
}

func New() *MyConfig{
	return &MyConfig{
		startPos : plate.New(640, 392, 1070, 429),
		shortOffset:  types.Offset{Y : 40, X : 0},
		longOffset:  types.Offset{Y : 160, X : 0},
	}
}

func (c *MyConfig) Plates() []plate.Plate{
	prevPos := c.startPos
	
	plates := []plate.Plate{c.startPos}

	for i := 0; i < 2; i++ {
		prevPos.Add(c.shortOffset)
		plates = append(plates, prevPos)
	}

	prevPos.Add(c.longOffset)

	for i := 0; i < 3; i++ {
		prevPos.Add(c.shortOffset)
		plates = append(plates, prevPos)
	}

	return plates
}