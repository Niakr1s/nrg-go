package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	log "github.com/sirupsen/logrus"
)

type Damage struct {
	reg *registry.Registry
}

func NewDamage(reg *registry.Registry) *Damage {
	return &Damage{reg: reg}
}

func (d *Damage) Step() {
	d.reg.RLock()
	defer d.reg.RUnlock()
	for i := 0; i < len(d.reg.Entities); i++ {
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
			dealDamage(lhs, rhs)
			dealDamage(rhs, lhs)
			rhs.Unlock()
		}
		lhs.Unlock()
	}
}

func dealDamage(attacker, defendant *entity.Entity) {
	dmgCs := attacker.GetComponents(component.DamageID)
	hpCs := defendant.GetComponents(component.HpID)
	if dmgCs == nil || hpCs == nil {
		return
	}
	dmg, hp := dmgCs[0].(component.Damage), hpCs[0].(component.HP)
	newHp := hp.Decrease(dmg.Dmg)
	log.Infof("dealt %d damage, old hp: %d, new hp: %d", dmg.Dmg, hp.Current, newHp.Current)
	defendant = defendant.SetComponents(newHp)
}
