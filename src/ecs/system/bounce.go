package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

// Bounce is a system, that bounces ground bodies with each other.
// It's a very naive implementation though.
type Bounce struct {
	reg *registry.Registry
}

func NewBounce(reg *registry.Registry) *Bounce {
	return &Bounce{reg: reg}
}

func (b *Bounce) Step() {
	b.reg.RLock()
	defer b.reg.RUnlock()

	for i := 0; i < len(b.reg.Entities); i++ {
		lhs := b.reg.Entities[i]
		lhs.Lock()
		lcs := lhs.GetComponents(component.PosID, component.ShapeID, component.GroundID)
		if lcs == nil {
			lhs.Unlock()
			continue
		}
		lhsPos := lcs[0].(component.Pos)
		lhsShape := lcs[1].(component.Shape)
		lhsGround := lcs[2].(component.Ground)
		for j := i + 1; j < len(b.reg.Entities); j++ {
			rhs := b.reg.Entities[j]
			rhs.Lock()
			rcs := rhs.GetComponents(component.PosID, component.ShapeID, component.GroundID)
			if rcs == nil {
				rhs.Unlock()
				continue
			}
			rhsPos := rcs[0].(component.Pos)
			rhsShape := rcs[1].(component.Shape)
			rhsGround := rcs[2].(component.Ground)
			if !lhsShape.Intersects(lhsPos, rhsPos, rhsShape) {
				rhs.Unlock()
				continue
			}
			lhsPos, rhsPos = lhsShape.BouncePos(lhsPos, rhsPos, lhsGround.Obstacle, rhsGround.Obstacle, rhsShape)
			lhs.SetComponents(lhsPos)
			rhs.SetComponents(rhsPos)
			rhs.Unlock()
		}
		lhs.Unlock()
	}
}
