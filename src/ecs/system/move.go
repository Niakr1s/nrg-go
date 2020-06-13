package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

type Move struct {
	reg *registry.Registry
}

func NewMove(reg *registry.Registry) *Move {
	return &Move{reg: reg}
}

func (m *Move) Step() {
	m.reg.RLock()
	defer m.reg.RUnlock()
	for _, e := range m.reg.Entities {
		e.Lock()
		if cs := e.GetComponents(component.VectorID, component.PosID, component.SpeedID); cs != nil {
			vec := cs[0].(*component.Vector)
			pos := cs[1].(*component.Pos)
			speed := cs[2].(*component.Speed)
			pos.Move(*vec, *speed)
		}
		e.Unlock()
	}
}
