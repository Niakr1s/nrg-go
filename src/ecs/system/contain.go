package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

// Contain is a system, that contains ground bodies inside board.
type Contain struct {
	reg            *registry.Registry
	boardW, boardH float64
}

func NewContain(reg *registry.Registry, boardW, boardH float64) *Contain {
	return &Contain{reg: reg, boardW: boardW, boardH: boardH}
}

func (c *Contain) Step() {
	c.reg.RLock()
	defer c.reg.RUnlock()
	for _, e := range c.reg.Entities {
		c.correctEntityInBoard(e)
	}
}

func (c *Contain) correctEntityInBoard(e *entity.Entity) {
	e.Lock()
	defer e.Unlock()
	cs := e.GetComponents(component.PosID, component.ShapeID, component.GroundID)
	if cs == nil {
		return
	}
	pos := cs[0].(component.Pos)
	shape := cs[1].(component.Shape)
	bound := shape.Bound(pos)
	pos, _ = correctPosBoard(pos, bound, c.boardW, c.boardH)
	e.SetComponents(pos)
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
