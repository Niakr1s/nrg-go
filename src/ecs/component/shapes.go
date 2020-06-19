package component

import (
	"math"

	log "github.com/sirupsen/logrus"
)

type Shape interface {
	Component
	Bound(center Pos) Bound
	Intersects(selfCenter, rhsCenter Pos, rhs Shape) bool
	BouncePos(selfCenter, rhsCenter Pos, selfIsObstacle, rhsIsObstacle bool, rhs Shape) (Pos, Pos)
	// OuterPointInDirectionDiff should return Pos on the outer contour of shape.
	OuterPointInDirectionDiff(vec Vector) Pos
}

type Bound struct {
	TopLeft, BotRight Pos
}

func NewBound(center Pos, w, h float64) Bound {
	left := center.X - w/2
	right := center.X + w/2
	top := center.Y - h/2
	bot := center.Y + h/2
	return Bound{TopLeft: NewPos(left, top), BotRight: NewPos(right, bot)}
}

func (b Bound) Outside(other Bound) bool {
	return b.TopLeft.X > other.BotRight.X || b.TopLeft.Y > other.BotRight.Y ||
		b.BotRight.X < other.TopLeft.X || b.BotRight.Y < other.TopLeft.Y
}

func (b Bound) Width() float64 {
	return b.BotRight.X - b.TopLeft.X
}

func (b Bound) Height() float64 {
	return b.BotRight.Y - b.TopLeft.Y
}

// Circle ..
type Circle struct {
	R float64
}

func NewCircle(r float64) Circle {
	return Circle{R: r}
}

func (c Circle) ID() ID {
	return ShapeID
}

func (c Circle) Bound(center Pos) Bound {
	return NewBound(center, c.R*2, c.R*2)
}

func (c Circle) Intersects(selfCenter, rhsCenter Pos, rhs Shape) bool {
	switch rhs := rhs.(type) {
	case Circle:
		dist := distance(selfCenter, rhsCenter)
		return dist < c.R+rhs.R

	default:
		log.Errorf("Circle.Intersects: got unknown rhs: %v", rhs)
		return false
	}
}

func (c Circle) BouncePos(selfCenter, rhsCenter Pos, selfIsObstacle, rhsIsObstacle bool, rhs Shape) (Pos, Pos) {
	if c.Intersects(selfCenter, rhsCenter, rhs) {
		switch rhs := rhs.(type) {
		case Circle:
			dist := distance(selfCenter, rhsCenter)
			diff := c.R + rhs.R - dist
			selfVec, rhsVec := NewVectorFromPos(rhsCenter, selfCenter), NewVectorFromPos(selfCenter, rhsCenter)
			// selfDist, rhsDist := diff*float64(selfSpeed)/float64(sumSpeed), diff*float64(rhsSpeed)/float64(sumSpeed)
			selfDist, rhsDist := diff*.5, diff*.5
			// we shouldn't move ground bodies
			if selfIsObstacle && !rhsIsObstacle {
				selfDist, rhsDist = 0, diff
			} else if !selfIsObstacle && rhsIsObstacle {
				selfDist, rhsDist = diff, 0
			}
			selfCenter, rhsCenter = selfCenter.Move(selfVec, Speed(selfDist)), rhsCenter.Move(rhsVec, Speed(rhsDist))

		default:
			log.Errorf("Circle.CorrectedPos: got unknown rhs: %v", rhs)
		}
	}

	return selfCenter, rhsCenter
}

func distance(lhs, rhs Pos) float64 {
	return math.Sqrt(math.Pow(lhs.X-rhs.X, 2) + math.Pow(lhs.Y-rhs.Y, 2))
}

// OuterPointInDirectionDiff returns Pos on the outer contour of shape.
func (c Circle) OuterPointInDirectionDiff(vec Vector) Pos {
	x, y := math.Cos(float64(vec))*c.R, math.Sin(float64(vec))*c.R
	return NewPos(x, y)
}
