package component

import "math"

// Vector is a vector in radians
// 0pi - right, 0.5pi - down, pi - left, 1.5pi - top
type Vector float64

func NewVector(vec float64) Vector {
	return Vector(vec)
}

func (v Vector) ID() ID {
	return VectorID
}

func (v Vector) IsLeft() bool {
	return math.Cos(float64(v)) <= 0
}

func (v Vector) IsRight() bool {
	return math.Cos(float64(v)) >= 0
}

func (v Vector) IsTop() bool {
	return math.Sin(float64(v)) <= 0
}

func (v Vector) IsBot() bool {
	return math.Sin(float64(v)) >= 0
}

type Speed float64

func NewSpeed(sp float64) Speed {
	return Speed(sp)
}

func (s Speed) ID() ID {
	return SpeedID
}

type Pos struct {
	X, Y float64
}

func NewPos(x, y float64) Pos {
	return Pos{x, y}
}

func (p Pos) ID() ID {
	return PosID
}

// Move returns Pos after move.
func (p Pos) Move(vec Vector, sp Speed) Pos {
	dx, dy := math.Cos(float64(vec))*float64(sp), math.Sin(float64(vec))*float64(sp)
	p.X += dx
	p.Y += dy
	return p
}
