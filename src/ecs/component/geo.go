package component

import "math"

// Vector is a vector in radians
// 0pi - right, 0.5pi - down, pi - left, 1.5pi - top
type Vector float64

func NewVector(vec float64) *Vector {
	res := Vector(vec)
	return &res
}

func (v *Vector) ID() ID {
	return VectorID
}

type Speed float64

func NewSpeed(sp float64) *Speed {
	res := Speed(sp)
	return &res
}

func (s *Speed) ID() ID {
	return SpeedID
}

type Pos struct {
	X, Y float64
}

func NewPos(x, y float64) *Pos {
	return &Pos{x, y}
}

func (p *Pos) ID() ID {
	return PosID
}

func (p *Pos) Move(vec Vector, sp Speed) {
	dx, dy := math.Cos(float64(vec))*float64(sp), math.Sin(float64(vec))*float64(sp)
	p.X += dx
	p.Y += dy
}
