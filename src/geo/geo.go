package geo

// Vector is a vector in radians
type Vector float64

type Distance float64

type Pos struct {
	X, Y float64
}

func NewPos(x, y float64) Pos {
	return Pos{x, y}
}

type Radius float64

func NewRadius(r float64) Radius {
	return Radius(r)
}
