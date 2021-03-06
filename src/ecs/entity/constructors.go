package entity

import (
	"math"

	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

func NewDefaultBullet(pos component.Pos, vec component.Vector, parent *Entity) *Entity {
	damage := component.NewDamage(10)
	res := NewEntity().
		SetComponents(
			component.NewCircle(20),
			pos,
			vec,
			component.NewSpeed(5),
			damage,
			NewParent(parent.ID),
		).
		SetTags(tag.Bullet)
	if cs := parent.GetComponents(component.FractionID); cs != nil {
		fraction := cs[0].(component.Fraction)
		res.SetComponents(fraction)
	}
	return res
}

func NewUser(pos component.Pos) *Entity {
	return NewEntity().
		SetComponents(
			component.NewCircle(50),
			pos,
			component.NewSpeed(10),
			component.NewGround(false),
			component.NewHP(100),
			component.FractionAlly,
		).
		SetTags(tag.User)
}

func NewEnemy(pos component.Pos) *Entity {
	return NewEntity().
		SetComponents(
			component.NewCircle(50),
			pos,
			component.NewSpeed(10),
			component.NewGround(true),
			component.NewHP(100),
			component.FractionEnemy,
		).
		SetTags(tag.Enemy)
}

func NewObstacle(pos component.Pos) *Entity {
	return NewEntity().
		SetComponents(
			component.NewCircle(50),
			pos,
			component.NewGround(true),
		)
}

func NewExplodeAnimation(pos component.Pos) *Entity {
	return NewEntity().
		SetComponents(
			pos,
			component.NewExplodeAnimation(),
		)
}

func NewEnemyWeaponWith4Guns() *component.AutoWeapon {
	enemyWeap := component.NewAutoWeapon(component.NewVector(0),
		component.NewVector(0),
		component.NewVector(0.5*math.Pi), component.NewVector(math.Pi), component.NewVector(1.5*math.Pi))
	enemyWeap.SetDirection(component.NewAutoWeaponDirection(1.5*math.Pi, component.NewVector(0.3*math.Pi)))
	return enemyWeap
}

func NewUserWeaponWith1Gun() *component.UserControlledWeapon {
	userWeap := component.NewUserControlledWeapon(component.NewVector(0))
	userWeap.SetDirection(component.NewUserControlledWeaponDirection(component.NewVector(1.5*math.Pi), component.NewVector(0.3*math.Pi)))
	return userWeap
}
