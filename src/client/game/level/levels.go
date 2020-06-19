package level

import (
	"math"

	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

func getInitLevelFuncs() []InitLevelFunc {
	res := []InitLevelFunc{
		level1(),
	}
	return res
}

func level1() InitLevelFunc {
	return func() *registry.Registry {
		reg := registry.NewRegistry()
		userWeap := component.NewUserControlledWeapon(component.NewVector(0))
		userWeap.SetDirection(component.NewUserControlledWeaponDirection(component.NewVector(1.5*math.Pi), component.NewVector(0.3*math.Pi)))

		enemyWeap := component.NewAutoWeapon(component.NewVector(0),
			component.NewVector(0),
			component.NewVector(0.5*math.Pi), component.NewVector(math.Pi), component.NewVector(1.5*math.Pi))
		enemyWeap.SetDirection(component.NewAutoWeaponDirection(1.5*math.Pi, component.NewVector(0.3*math.Pi)))

		reg.AddEntity(entity.NewUser(component.NewPos(500, 500)).SetComponents(userWeap))
		reg.AddEntity(entity.NewEnemy(component.NewPos(200, 200)).SetComponents(enemyWeap))
		reg.AddEntity(entity.NewObstacle(component.NewPos(700, 200)))
		return reg
	}
}
