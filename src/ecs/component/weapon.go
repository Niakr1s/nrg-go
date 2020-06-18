package component

import (
	"sync"
	"time"

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

	ReloadDuration time.Duration

	readyMutex sync.Mutex
	ready      bool
}

// NewAutoWeapon constructs Weapon with default reload duration and constant Direction.
func NewAutoWeapon(gunsDirDiffs ...Vector) *AutoWeapon {
	res := &AutoWeapon{
		Dir:            NewVector(0),
		Guns:           make([]*Gun, len(gunsDirDiffs)),
		ReloadDuration: config.ReloadDuration,
		ready:          true,
	}
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

func (w *AutoWeapon) SetReloadDuration(d time.Duration) *AutoWeapon {
	w.ReloadDuration = d
	return w
}

func (w *AutoWeapon) IsReady() bool {
	w.readyMutex.Lock()
	defer w.readyMutex.Unlock()
	return w.ready
}

// Fire gets directions of all guns and starts reload timer.
func (w *AutoWeapon) Fire() bool {
	if !w.IsReady() {
		return false
	}
	w.startReloading()
	return true
}

func (w *AutoWeapon) GetGunDirs() []Vector {
	res := make([]Vector, len(w.Guns))
	for i, g := range w.Guns {
		res[i] = g.DirectionDiff.Sum(w.Dir.Direction())
	}
	return res
}

func (w *AutoWeapon) startReloading() {
	w.ready = false
	go func() {
		<-time.After(w.ReloadDuration)
		w.readyMutex.Lock()
		defer w.readyMutex.Unlock()
		w.ready = true
	}()
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
		return false
	}
	return w.AutoWeapon.Fire()
}

func (w *UserControlledWeapon) SetAutoAttack(attack bool) {
	w.autoAttackMutex.Lock()
	defer w.autoAttackMutex.Unlock()
	w.autoAttack = attack
}
