package component

import (
	"time"

	log "github.com/sirupsen/logrus"
)

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

type UserControlledWeaponDirection struct {
	WeaponDirection

	defaultAngleSpeedPerSec Vector
}

func NewUserControlledWeaponDirection(startVec, angleSpeedPerSec Vector) UserControlledWeaponDirection {
	return UserControlledWeaponDirection{
		WeaponDirection:         NewAutoWeaponDirection(startVec, NewVector(0)),
		defaultAngleSpeedPerSec: angleSpeedPerSec,
	}
}

type WeaponDir int

const (
	WeaponDirLeft WeaponDir = iota - 1
	WeaponDirZero
	WeaponDirRight
)

func (uwd UserControlledWeaponDirection) NewDirection(dir WeaponDir) UserControlledWeaponDirection {
	oldDir := uwd.WeaponDirection.Direction()
	angleSpeed := NewVector(float64(uwd.defaultAngleSpeedPerSec) * float64(dir))
	log.Tracef("UserControlledWeaponDirection.SetDirection(): direction changed: %v, oldDir=%v, angleSpeed=%v", dir, oldDir, angleSpeed)
	uwd.WeaponDirection = NewAutoWeaponDirection(oldDir, angleSpeed)
	return uwd
}
