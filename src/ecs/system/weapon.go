package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

type Weapon struct{}

func NewWeapon() *Weapon {
	return &Weapon{}
}

func (w *Weapon) Step(reg *registry.Registry) {
	bullets := w.spawnBullets(reg)
	w.addBulletsToRegistry(reg, bullets)
}

func (w *Weapon) spawnBullets(reg *registry.Registry) []*entity.Entity {
	reg.RLock()
	defer reg.RUnlock()

	bullets := []*entity.Entity{}
	for _, e := range reg.Entities {
		e.RLock()
		cs := e.GetComponents(component.PosID, component.ShapeID, component.WeaponID)
		if cs == nil {
			e.RUnlock()
			continue
		}
		pos := cs[0].(component.Pos)
		shape := cs[1].(component.Shape)
		weap := cs[2].(component.Weapon)

		if weap.Fire() {
			bulletDirections := weap.GetGunDirs()
			for _, bdir := range bulletDirections {
				// creating bullet in the center of entity
				bullet := entity.NewDefaultBullet(pos, bdir, e.ID)
				bulletShape := bullet.GetComponents(component.ShapeID)[0].(component.Shape)
				moveDiff := shape.OuterPointInDirectionDiff(bdir).Sum(bulletShape.OuterPointInDirectionDiff(bdir))

				newPos := pos.Sum(moveDiff)
				bullet.SetComponents(newPos)
				// log.Tracef("spawned bullet at %v", newPos)

				bullets = append(bullets, bullet)
			}
		}

		e.RUnlock()
	}

	return bullets
}

func (w *Weapon) addBulletsToRegistry(reg *registry.Registry, bullets []*entity.Entity) {
	if len(bullets) == 0 {
		return
	}
	reg.Lock()
	defer reg.Unlock()
	for _, b := range bullets {
		reg.AddEntity(b)
	}
}
