package component

import "time"

// WeaponDirection must return valid direction each time it is called.
type WeaponDirection interface {
	Direction() Vector
}

type AutoWeaponDirection struct {
	vec              Vector
	angleSpeedPerSec Vector
	createdAt        time.Time
}

func NewAutoWeaponDirection(startVec, angleSpeedPerSec Vector) AutoWeaponDirection {
	return AutoWeaponDirection{
		vec:              startVec,
		angleSpeedPerSec: angleSpeedPerSec,
		createdAt:        time.Now(),
	}
}

func (wd AutoWeaponDirection) Direction() Vector {
	elapsed := time.Now().Sub(wd.createdAt)
	angleDiff := elapsed.Seconds() * float64(wd.angleSpeedPerSec)
	return wd.vec.Sum(NewVector(angleDiff))
}
