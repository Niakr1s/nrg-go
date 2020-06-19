package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
	log "github.com/sirupsen/logrus"
)

type Damage struct{}

func NewDamage() *Damage {
	return &Damage{}
}

func (d *Damage) Step(reg *registry.Registry) {
	destroyedPoses := d.dealDamageAll(reg)
	d.addExplosions(reg, destroyedPoses)
}

func (d *Damage) addExplosions(reg *registry.Registry, poses []component.Pos) {
	reg.Lock()
	defer reg.Unlock()
	for _, explosionPos := range poses {
		reg.AddEntity(entity.NewExplodeAnimation(explosionPos))
	}
}

// dealDamageAll iterates through whole registry and deals damage.
// Returns poses of destroyed entities.
func (d *Damage) dealDamageAll(reg *registry.Registry) []component.Pos {
	destroyedPoses := []component.Pos{}
	reg.RLock()
	defer reg.RUnlock()
	for i := range reg.Entities {
		lhs := reg.Entities[i]
		lhs.Lock()
		lcs := lhs.GetComponents(component.PosID, component.ShapeID)
		if lcs == nil {
			lhs.Unlock()
			continue
		}
		lPos := lcs[0].(component.Pos)
		lShape := lcs[1].(component.Shape)
		for j := i + 1; j < len(reg.Entities); j++ {
			rhs := reg.Entities[j]
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
	if dmg.AlliedTags.IsAllyWith(defendant.Tags) {
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
