package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
	log "github.com/sirupsen/logrus"
)

type Damage struct {
	reg *registry.Registry
}

func NewDamage(reg *registry.Registry) *Damage {
	return &Damage{reg: reg}
}

func (d *Damage) Step() {
	destroyedPoses := d.dealDamageAll()
	d.addExplosions(destroyedPoses)
}

func (d *Damage) addExplosions(poses []component.Pos) {
	d.reg.Lock()
	defer d.reg.Unlock()
	for _, explosionPos := range poses {
		d.reg.AddEntity(entity.NewExplodeAnimation(explosionPos))
	}
}

// dealDamageAll iterates through whole registry and deals damage.
// Returns poses of destroyed entities.
func (d *Damage) dealDamageAll() []component.Pos {
	destroyedPoses := []component.Pos{}
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
			_, destroyed := dealDamage(lhs, rhs)
			if destroyed {
				destroyedPoses = append(destroyedPoses, rPos)
			}
			_, destroyed = dealDamage(rhs, lhs)
			if destroyed {
				destroyedPoses = append(destroyedPoses, lPos)
			}
			rhs.Unlock()
		}
		lhs.Unlock()
	}
	return destroyedPoses
}

func dealDamage(attacker, defendant *entity.Entity) (damaged bool, destroyed bool) {
	dmgCs := attacker.GetComponents(component.DamageID)
	hpCs := defendant.GetComponents(component.HpID)
	if dmgCs == nil || hpCs == nil {
		return
	}
	dmg, hp := dmgCs[0].(component.Damage), hpCs[0].(component.HP)
	if parentCs := attacker.GetComponents(component.ParentID); parentCs != nil &&
		parentCs[0].(entity.Parent).Parent == defendant.ID {
		return
	}
	newHp := hp.Decrease(dmg.Dmg)
	damaged = true
	log.Infof("dealt %d damage, old hp: %d, new hp: %d", dmg.Dmg, hp.Current, newHp.Current)
	defendant = defendant.SetComponents(newHp)
	if newHp.IsDead() {
		destroyed = true
		defendant.SetTags(tag.Destroyed)
	}
	return
}
