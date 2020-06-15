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
