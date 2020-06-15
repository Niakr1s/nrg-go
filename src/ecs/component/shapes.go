package component

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/img"
	log "github.com/sirupsen/logrus"
)

type Shape interface {
	Component
	Draw(board *ebiten.Image, pos Pos)
	Bound(center Pos) Bound
	Intersects(selfCenter, rhsCenter Pos, rhs Shape) bool
	BouncePos(selfCenter, rhsCenter Pos, selfIsObstacle, rhsIsObstacle bool, rhs Shape) (Pos, Pos)
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

// Circle ..
type Circle struct {
	R float64

	image *ebiten.Image
}

func NewCircle(r float64, path string) Circle {
	image, err := img.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return Circle{R: r, image: image}
}

func (c Circle) ID() ID {
	return ShapeID
}

func (c Circle) Draw(board *ebiten.Image, pos Pos) {
	op := &ebiten.DrawImageOptions{}
	w, h := c.image.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	scale := float64(c.R*2) / float64(w)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(pos.X, pos.Y)

	board.DrawImage(c.image, op)
}

func (c Circle) Bound(center Pos) Bound {
	return NewBound(center, c.R*2, c.R*2)
}

func (c Circle) Intersects(selfCenter, rhsCenter Pos, rhs Shape) bool {
	switch rhs := rhs.(type) {
	case Circle:
		dist := distance(selfCenter, rhsCenter)
		return dist <= c.R+rhs.R

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
