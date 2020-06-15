package component

type HP struct {
	Current int
}

func NewHP(maxHP int) HP {
	return HP{Current: maxHP}
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

type Damage struct {
	Dmg int
}

func NewDamage(dmg int) Damage {
	return Damage{Dmg: dmg}
}

func (d Damage) ID() ID {
	return DamageID
}
