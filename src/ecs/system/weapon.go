package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	log "github.com/sirupsen/logrus"
)

type Weapon struct {
	reg *registry.Registry
}

func NewWeapon(reg *registry.Registry) *Weapon {
	return &Weapon{reg: reg}
}

func (w *Weapon) Step() {
	bullets := w.spawnBullets()
	w.addBulletsToRegistry(bullets)
}

func (w *Weapon) spawnBullets() []*entity.Entity {
	w.reg.RLock()
	defer w.reg.RUnlock()

	bullets := []*entity.Entity{}
	for _, e := range w.reg.Entities {
		e.RLock()
		cs := e.GetComponents(component.PosID, component.ShapeID, component.WeaponID)
		if cs == nil {
			e.RUnlock()
			continue
		}
		pos := cs[0].(component.Pos)
		shape := cs[1].(component.Shape)
		weap := cs[2].(*component.Weapon)

		bulletDirections := weap.Fire()
		for _, bdir := range bulletDirections {
			// creating bullet in the center of entity
			bullet := entity.NewDefaultBullet(pos, bdir)
			bulletShape := bullet.GetComponents(component.ShapeID)[0].(component.Shape)
			moveDiff := shape.OuterPointInDirectionDiff(bdir).Sum(bulletShape.OuterPointInDirectionDiff(bdir))

			newPos := pos.Sum(moveDiff)
			bullet.SetComponents(newPos)
			log.Tracef("spawned bullet at %v", newPos)

			bullets = append(bullets, bullet)
		}

		e.RUnlock()
	}

	return bullets
}

func (w *Weapon) addBulletsToRegistry(bullets []*entity.Entity) {
	if len(bullets) == 0 {
		return
	}
	w.reg.Lock()
	defer w.reg.Unlock()
	lenBefore := len(w.reg.Entities)
	for _, b := range bullets {
		w.reg.AddEntity(b)
	}
	lenAfter := len(w.reg.Entities)
	log.Tracef("Weapon.addBulletsToRegistry: added %d bullets, len(entities): %d -> %d",
		len(bullets), lenBefore, lenAfter)
}
