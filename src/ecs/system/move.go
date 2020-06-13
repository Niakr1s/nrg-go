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
}

func (m *Move) moveOneEntity(e *entity.Entity) {
	e.Lock()
	defer e.Unlock()
	if cs := e.GetComponents(component.VectorID, component.PosID, component.SpeedID, component.ShapeID); cs != nil {
		vec := cs[0].(component.Vector)
		pos := cs[1].(component.Pos)
		speed := cs[2].(component.Speed)
		shape := cs[3].(component.Shape)
		bound := shape.Bound(pos)
		if bound.TopLeft.X < 1 && vec.IsLeft() ||
			bound.BotRight.X > m.boardW-1 && vec.IsRight() ||
			bound.TopLeft.Y < 1 && vec.IsTop() ||
			bound.BotRight.Y > m.boardH-1 && vec.IsBot() {
			return
		}
		e.SetComponents(pos.Move(vec, speed))
	}
}
