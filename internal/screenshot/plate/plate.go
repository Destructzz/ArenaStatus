package plate

import "ArenaStatus/internal/types"

type Plate struct {
	X0, Y0, X1, Y1 int
}

func New(x0, y0, x1, y1 int) Plate {
	return Plate{
		X0: x0,
		Y0: y0,
		X1: x1,
		Y1: y1,
	}
}

func (p *Plate) Add(offset types.Offset) {
	p.X1 += offset.X
	p.X0 += offset.X
	p.Y0 += offset.Y
	p.Y1 += offset.Y
}