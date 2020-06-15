package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

type Move struct {
	reg            *registry.Registry
	boardW, boardH float64
}

func NewMove(reg *registry.Registry, boardW, boardH float64) *Move {
	return &Move{reg: reg, boardW: boardW, boardH: boardH}
}

func (m *Move) Step() {
	m.reg.RLock()
	defer m.reg.RUnlock()
	for _, e := range m.reg.Entities {
		m.moveOneEntity(e)
	}
}
func (m *Move) correctPos() {
	for i := 0; i < len(m.reg.Entities); i++ {
		lhs := m.reg.Entities[i]
		lhs.Lock()
		lcs := lhs.GetComponents(component.PosID, component.ShapeID)
		if lcs == nil || !lhs.HasTags(tag.GroundID) {
			lhs.Unlock()
			continue
		}
		lhsPos := lcs[0].(component.Pos)
		lhsShape := lcs[1].(component.Shape)
		for j := i + 1; j < len(m.reg.Entities); j++ {
			rhs := m.reg.Entities[j]
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

func (m *Move) moveOneEntity(e *entity.Entity) {
	e.Lock()
	defer e.Unlock()
	cs := e.GetComponents(component.VectorID, component.PosID, component.SpeedID)
	if cs == nil {
		return
	}
	vec := cs[0].(component.Vector)
	pos := cs[1].(component.Pos)
	speed := cs[2].(component.Speed)
	// no speed - no move
	if speed <= 0 {
		e.RemoveComponents(component.SpeedID)
		return
	}
	e.SetComponents(pos.Move(vec, speed))
}
