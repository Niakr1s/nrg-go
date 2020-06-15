package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

type Destroy struct {
	reg            *registry.Registry
	boardW, boardH float64
}

func NewDestroy(reg *registry.Registry, boardW, boardH float64) *Destroy {
	return &Destroy{reg: reg, boardW: boardW, boardH: boardH}
}

func (d *Destroy) Step() {
	d.destroyBulletsOutOfBoard()
	d.destroyBulletsContactedWithGroundBodies()
}

func (d *Destroy) destroyBulletsContactedWithGroundBodies() {
	d.reg.RLock()
	defer d.reg.RUnlock()
	for i := range d.reg.Entities {
		lhs := d.reg.Entities[i]
		lhs.Lock()
		lcs := lhs.GetComponents(component.PosID, component.ShapeID)
		if lcs == nil {
			lhs.Unlock()
			continue
		}
		lPos := lcs[0].(component.Pos)
		lShape := lcs[1].(component.Shape)
		for j := i + 1; j < len(d.reg.Entities); j++ {
			rhs := d.reg.Entities[j]
			rhs.Lock()
			rcs := rhs.GetComponents(component.PosID, component.ShapeID)
			if rcs == nil {
				rhs.Unlock()
				continue
			}
			rPos := rcs[0].(component.Pos)
			rShape := rcs[1].(component.Shape)
			if !lShape.Intersects(lPos, rPos, rShape) {
				rhs.Unlock()
				continue
			}
			checkAndDestroyBulletContactedWithBody(lhs, rhs)
			checkAndDestroyBulletContactedWithBody(rhs, lhs)
			rhs.Unlock()
		}
		lhs.Unlock()
	}
}

func checkAndDestroyBulletContactedWithBody(bullet, body *entity.Entity) {
	if !bullet.HasTags(tag.Bullet) || body.GetComponents(component.GroundID) == nil {
		return
	}
	bullet.SetTags(tag.Destroyed)
}

func (d *Destroy) destroyBulletsOutOfBoard() {
	boardBound := component.Bound{TopLeft: component.NewPos(0, 0), BotRight: component.NewPos(d.boardW, d.boardH)}
	d.reg.RLock()
	defer d.reg.RUnlock()
	for _, e := range d.reg.Entities {
		e.Lock()
		cs := e.GetComponents(component.PosID, component.ShapeID)
		if cs == nil || !e.HasTags(tag.Bullet) {
			e.Unlock()
			continue
		}
		pos, shape := cs[0].(component.Pos), cs[1].(component.Shape)
		bound := shape.Bound(pos)
		if bound.Outside(boardBound) {
			e.SetTags(tag.Destroyed)
		}
		e.Unlock()
	}
}
