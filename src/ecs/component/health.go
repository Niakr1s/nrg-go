package component

import "github.com/niakr1s/nrg-go/src/ecs/tag"

type HP struct {
	MaxHP   int
	Current int
}

func NewHP(maxHP int) HP {
	return HP{Current: maxHP, MaxHP: maxHP}
}

func (hp HP) ID() ID {
	return HpID
}

func (hp HP) Decrease(amount int) HP {
	hp.Current -= amount
	return hp
}

func (hp HP) IsDead() bool {
	return hp.Current <= 0
}

func (hp HP) Percent() float64 {
	return float64(hp.Current) / float64(hp.MaxHP)
}

type Damage struct {
	Dmg        int
	AlliedTags tag.Tags
}

func NewDamage(dmg int) Damage {
	return Damage{Dmg: dmg, AlliedTags: make(tag.Tags)}
}

func (d Damage) ID() ID {
	return DamageID
}
