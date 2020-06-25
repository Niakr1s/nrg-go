package component

import (
	"sync"

	"github.com/niakr1s/nrg-go/src/config"
)

type Weapon interface {
	Fire() bool
	GetGunDirs() []Vector
	WeaponDir() WeaponDirection
	SetDirection(WeaponDirection)
}

type AutoWeapon struct {
	Dir  WeaponDirection
	Guns []*Gun

	// ReloadDuration in frames, 1 sec = 60 frames
	ReloadDuration int
	// Reloaded indicates how much reloaded is weapon
	Reloaded int
}

// NewAutoWeapon constructs Weapon with default reload duration and constant Direction.
func NewAutoWeapon(gunsDirDiffs ...Vector) *AutoWeapon {
	res := &AutoWeapon{
		Dir:  NewVector(0),
		Guns: make([]*Gun, len(gunsDirDiffs)),
	}
	res.SetReloadDuration(config.ReloadDuration)
	for i, diff := range gunsDirDiffs {
		res.Guns[i] = NewGun(res, diff)
	}
	return res
}

func (w *AutoWeapon) SetDirection(dir WeaponDirection) {
	w.Dir = dir
}

func (w *AutoWeapon) ID() ID {
	return WeaponID
}

func (w *AutoWeapon) WeaponDir() WeaponDirection {
	return w.Dir
}

func (w *AutoWeapon) SetReloadDuration(d int) *AutoWeapon {
	w.ReloadDuration = d
	w.Reloaded = d // we want weapon to be reloaded initially
	return w
}

func (w *AutoWeapon) IsReady() bool {
	return w.Reloaded >= w.ReloadDuration
}

// Fire gets directions of all guns and starts reload timer.
func (w *AutoWeapon) Fire() bool {
	w.IncrementReloaded()
	if !w.IsReady() {
		return false
	}
	w.Reloaded = 0
	return true
}

// IncrementReloaded increments reloaded, it must be called each Fire() call, even in inheritors.
func (w *AutoWeapon) IncrementReloaded() {
	w.Reloaded++
}

func (w *AutoWeapon) GetGunDirs() []Vector {
	res := make([]Vector, len(w.Guns))
	for i, g := range w.Guns {
		res[i] = g.DirectionDiff.Sum(w.Dir.Direction())
	}
	return res
}

type Gun struct {
	DirectionDiff Vector
}

func NewGun(weap *AutoWeapon, diff Vector) *Gun {
	return &Gun{DirectionDiff: diff}
}

type UserControlledWeapon struct {
	*AutoWeapon

	autoAttackMutex sync.Mutex
	autoAttack      bool
}

func NewUserControlledWeapon(gunsDirDiffs ...Vector) *UserControlledWeapon {
	res := &UserControlledWeapon{
		AutoWeapon: NewAutoWeapon(gunsDirDiffs...),
	}
	return res
}

func (w *UserControlledWeapon) Fire() bool {
	w.autoAttackMutex.Lock()
	defer w.autoAttackMutex.Unlock()
	if !w.autoAttack {
		// we still want to continue reload weapon
		w.AutoWeapon.IncrementReloaded()
		return false
	}
	return w.AutoWeapon.Fire()
}

func (w *UserControlledWeapon) SetAutoAttack(attack bool) {
	w.autoAttackMutex.Lock()
	defer w.autoAttackMutex.Unlock()
	w.autoAttack = attack
}
