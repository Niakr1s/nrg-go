package component

type Fraction int

const (
	FractionEnemy Fraction = iota
	FractionAlly
)

func (f Fraction) IsAllyWith(rhs Fraction) bool {
	return f == rhs
}

func (f Fraction) ID() ID {
	return FractionID
}
