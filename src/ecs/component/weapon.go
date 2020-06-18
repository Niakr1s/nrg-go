package component

import (
	"time"

	"github.com/niakr1s/nrg-go/src/config"
)

type Weapon struct {
	Direction Vector
	Guns      []*Gun

	ReloadDuration time.Duration
	reloadDone     chan struct{}
	ready          bool
}

// NewWeapon constructs Weapon with default reload duration.
func NewWeapon(dir Vector, gunsDirDiffs ...Vector) *Weapon {
	res := &Weapon{
		Direction:      dir,
		Guns:           make([]*Gun, len(gunsDirDiffs)),
		ReloadDuration: config.ReloadDuration,
		reloadDone:     make(chan struct{}, 1),
		ready:          true,
	}
	for i, diff := range gunsDirDiffs {
		res.Guns[i] = NewGun(res, diff)
	}
	return res
}

func (w *Weapon) ID() ID {
	return WeaponID
}

func (w *Weapon) SetReloadDuration(d time.Duration) *Weapon {
	w.ReloadDuration = d
	return w
}

func (w *Weapon) IsReady() bool {
	select {
	case <-w.reloadDone:
		w.ready = true
	default:
	}
	return w.ready
}

// Fire gets directions of all guns and starts reload timer.
func (w *Weapon) Fire() []Vector {
	if !w.IsReady() {
		return []Vector{}
	}
	res := make([]Vector, len(w.Guns))
	for i, g := range w.Guns {
		res[i] = g.DirectionDiff.Sum(w.Direction)
	}
	w.startReloading()
	return res
}

func (w *Weapon) startReloading() {
	w.ready = false
	go func() {
		<-time.After(w.ReloadDuration)
		w.reloadDone <- struct{}{}
	}()
}

type Gun struct {
	DirectionDiff Vector
}

func NewGun(weap *Weapon, diff Vector) *Gun {
	return &Gun{DirectionDiff: diff}
}
