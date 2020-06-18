package component

// WeaponDirection must return valid direction each time it is called.
type WeaponDirection interface {
	Direction() Vector
}
