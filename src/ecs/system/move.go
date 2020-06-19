package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

type Move struct {
	boardW, boardH float64
}

func NewMove(boardW, boardH float64) *Move {
	return &Move{boardW: boardW, boardH: boardH}
}

func (m *Move) Step(reg *registry.Registry) {
	reg.RLock()
	defer reg.RUnlock()
	for _, e := range reg.Entities {
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
