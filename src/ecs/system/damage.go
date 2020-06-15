package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

type Damage struct {
	reg *registry.Registry
}

func NewDamage(reg *registry.Registry) *Damage {
	return &Damage{reg: reg}
}

func (d *Damage) Step() {
	d.reg.Lock()
	d.reg.Unlock()
	for i := range d.reg.Entities {
		lhs := d.reg.Entities[i]
		lhs.Lock()
		lcs := lhs.GetComponents(component.PosID, component.ShapeID)
		lDamageCs := lhs.GetComponents(component.DamageID)
		lHpCs := lhs.GetComponents(component.HpID)
		if lcs == nil || (lDamageCs == nil && lHpCs == nil) {
			lhs.Unlock()
			continue
		}
		lPos := lcs[0].(component.Pos)
		lShape := lcs[1].(component.Shape)
		for j := i + 1; j < len(d.reg.Entities); j++ {
			rhs := d.reg.Entities[j]
			rhs.Lock()
			rcs := rhs.GetComponents(component.PosID, component.ShapeID)
			rDamageCs := rhs.GetComponents(component.DamageID)
			rHpCs := rhs.GetComponents(component.HpID)
			if rcs == nil || (rDamageCs == nil && rHpCs == nil) {
				rhs.Unlock()
				continue
			}
			rPos := rcs[0].(component.Pos)
			rShape := rcs[1].(component.Shape)
			if !lShape.Intersects(lPos, rPos, rShape) {
				rhs.Unlock()
				continue
			}
			if lHpCs != nil && rDamageCs != nil {
				lHp, rDmg := lHpCs[0].(component.HP), rDamageCs[0].(component.Damage)
				lhs.SetComponents(lHp.Decrease(rDmg.Dmg))
			}
			if rHpCs != nil && lDamageCs != nil {
				rHp, lDmg := lHpCs[0].(component.HP), rDamageCs[0].(component.Damage)
				rhs.SetComponents(rHp.Decrease(lDmg.Dmg))
			}
			rhs.Unlock()
		}
		lhs.Unlock()
	}
}
