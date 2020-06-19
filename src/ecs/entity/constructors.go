package entity

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

func NewDefaultBullet(pos component.Pos, vec component.Vector, parent EntityID) *Entity {
	return NewEntity().
		SetComponents(
			component.NewCircle(20),
			pos,
			vec,
			component.NewSpeed(5),
			component.NewDamage(10),
			NewParent(parent),
		).
		SetTags(tag.Bullet)
}

func NewUser(pos component.Pos) *Entity {
	return NewEntity().
		SetComponents(
			component.NewCircle(50),
			pos,
			component.NewSpeed(10),
			component.NewGround(false),
			component.NewHP(100),
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
