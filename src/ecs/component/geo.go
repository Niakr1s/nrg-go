package component

import "math"

// Vector is a vector in radians
// 0pi - right, 0.5pi - down, pi - left, 1.5pi - top
type Vector float64

func NewVector(vec float64) Vector {
	return Vector(vec)
}

func NewVectorFromPos(pos1, pos2 Pos) Vector {
	dx, dy := pos2.X-pos1.X, pos2.Y-pos1.Y
	a := math.Atan2(dy, dx)
	return Vector(a)
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

func (v Vector) Sum(rhs Vector) Vector {
	return NewVector(float64(rhs) + float64(v))
}

func (v Vector) Cos() float64 {
	return math.Cos(float64(v))
}

func (v Vector) Sin() float64 {
	return math.Sin(float64(v))
}

func (v Vector) Mirrored() Vector {
	diffX, diffY := v.Cos(), v.Sin()
	return NewVectorFromPos(NewPos(0, 0), NewPos(-diffX, -diffY))
}

func (v Vector) Direction() Vector {
	return v
}

// default vectors
var (
	RightVec Vector
	BotVec   Vector
	LeftVec  Vector
	TopVec   Vector
)

func init() {
	RightVec = NewVector(0.0 * math.Pi)
	BotVec = NewVector(0.5 * math.Pi)
	LeftVec = NewVector(1.0 * math.Pi)
	TopVec = NewVector(1.5 * math.Pi)
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

// MoveDist returns Pos after move.
func (p Pos) MoveDist(vec Vector, dist float64) Pos {
	dx, dy := math.Cos(float64(vec))*dist, math.Sin(float64(vec))*dist
	p.X += dx
	p.Y += dy
	return p
}

func (p Pos) Sum(rhs Pos) Pos {
	rhs.X += p.X
	rhs.Y += p.Y
	return rhs
}
