package level

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

func getInitLevelFuncs() []InitLevelFunc {
	res := []InitLevelFunc{
		level1(),
		level2(),
	}
	return res
}

func level1() InitLevelFunc {
	return func() *registry.Registry {
		reg := registry.NewRegistry()

		reg.AddEntity(entity.NewUser(component.NewPos(500, 800)).SetComponents(entity.NewUserWeaponWith1Gun()))
		reg.AddEntity(entity.NewEnemy(component.NewPos(500, 200)).SetComponents(entity.NewEnemyWeaponWith4Guns()))
		return reg
	}
}

func level2() InitLevelFunc {
	return func() *registry.Registry {
		reg := registry.NewRegistry()

		reg.AddEntity(entity.NewUser(component.NewPos(500, 800)).SetComponents(entity.NewUserWeaponWith1Gun()))
		reg.AddEntity(entity.NewEnemy(component.NewPos(300, 200)).SetComponents(entity.NewEnemyWeaponWith4Guns()))
		reg.AddEntity(entity.NewEnemy(component.NewPos(700, 200)).SetComponents(entity.NewEnemyWeaponWith4Guns()))
		return reg
	}
}
