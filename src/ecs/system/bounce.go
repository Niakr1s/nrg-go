package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
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
		lcs := lhs.GetComponents(component.PosID, component.ShapeID)
		if lcs == nil || !lhs.HasTags(tag.GroundID) {
			lhs.Unlock()
			continue
		}
		lhsPos := lcs[0].(component.Pos)
		lhsShape := lcs[1].(component.Shape)
		for j := i + 1; j < len(b.reg.Entities); j++ {
			rhs := b.reg.Entities[j]
			rhs.Lock()
			rcs := rhs.GetComponents(component.PosID, component.ShapeID)
			if rcs == nil || !rhs.HasTags(tag.GroundID) {
				rhs.Unlock()
				continue
			}
			rhsPos := rcs[0].(component.Pos)
			rhsShape := rcs[1].(component.Shape)
			if !lhsShape.Intersects(lhsPos, rhsPos, rhsShape) {
				rhs.Unlock()
				continue
			}
			lhsPos, rhsPos = lhsShape.CorrectedPos(lhsPos, rhsPos, lhs.HasTags(tag.GroundID), rhs.HasTags(tag.GroundID), rhsShape)
			lhs.SetComponents(lhsPos)
			rhs.SetComponents(rhsPos)
			rhs.Unlock()
		}
		lhs.Unlock()
	}
}
