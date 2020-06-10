package geo

import "math"

// Vector is a vector in radians
// 0pi - right, 0.5pi - down, pi - left, 1.5pi - top
type Vector float64

func (v Vector) Vector() Vector {
	return v
}

func NewVector(vec float64) Vector {
	return Vector(vec)
}

type Distance float64

type Pos struct {
	X, Y float64
}

func NewPos(x, y float64) Pos {
	return Pos{x, y}
}

func (p *Pos) Move(vec Vector, dist Distance) {
	dx, dy := math.Cos(float64(vec))*float64(dist), math.Sin(float64(vec))*float64(dist)
	p.X += dx
	p.Y += dy
}

type Radius float64

func NewRadius(r float64) Radius {
	return Radius(r)
}
