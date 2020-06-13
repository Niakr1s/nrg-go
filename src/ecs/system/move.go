package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
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
	m.correctPos()
}
func (m *Move) correctPos() {
	for i := 0; i < len(m.reg.Entities); i++ {
		lhs := m.reg.Entities[i]
		lcs := lhs.GetComponents(component.VectorID, component.PosID, component.SpeedID, component.ShapeID)
		if lcs == nil {
			continue
		}
		// lhsVec := lcs[0].(component.Vector)
		lhsPos := lcs[1].(component.Pos)
		// lhsSpeed := lcs[2].(component.Speed)
		lhsShape := lcs[3].(component.Shape)
		for j := i + 1; j < len(m.reg.Entities); j++ {
			rhs := m.reg.Entities[j]
			rcs := rhs.GetComponents(component.VectorID, component.PosID, component.SpeedID, component.ShapeID)
			if rcs == nil {
				continue
			}
			// rhsVec := rcs[0].(component.Vector)
			rhsPos := rcs[1].(component.Pos)
			// rhsSpeed := rcs[2].(component.Speed)
			rhsShape := rcs[3].(component.Shape)
			if !lhsShape.Intersects(lhsPos, rhsPos, rhsShape) {
				continue
			}
		}
	}
}

func (m *Move) moveOneEntity(e *entity.Entity) {
	e.Lock()
	defer e.Unlock()
	cs := e.GetComponents(component.VectorID, component.PosID, component.SpeedID, component.ShapeID)
	if cs == nil {
		return
	}
	vec := cs[0].(component.Vector)
	pos := cs[1].(component.Pos)
	speed := cs[2].(component.Speed)
	shape := cs[3].(component.Shape)
	// no speed - no move
	if speed <= 0 {
		e.RemoveComponents(component.SpeedID)
		return
	}
	nextPos := pos.Move(vec, speed)
	bound := shape.Bound(nextPos)
	nextPos, _ = correctPosBoard(nextPos, bound, m.boardW, m.boardH)
	e.SetComponents(nextPos)
}

func correctPosBoard(pos component.Pos, bound component.Bound, boardW, boardH float64) (component.Pos, bool) {
	changed := false
	if diff := 0 - bound.TopLeft.X; diff > 0 {
		pos.X += diff
		changed = true
	}
	if diff := 0 - bound.TopLeft.Y; diff > 0 {
		changed = true
		pos.Y += diff
	}
	if diff := boardW - bound.BotRight.X; diff < 0 {
		changed = true
		pos.X += diff
	}
	if diff := boardH - bound.BotRight.Y; diff < 0 {
		changed = true
		pos.Y += diff
	}
	return pos, changed
}
