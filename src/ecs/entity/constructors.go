package entity

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

func NewDefaultBullet(pos component.Pos, vec component.Vector) *Entity {
	return NewEntity().
		SetComponents(
			component.NewCircle(20),
			pos,
			vec,
			component.NewSpeed(5),
			component.NewDamage(10),
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
		SetTags(tag.User, tag.Player)
}