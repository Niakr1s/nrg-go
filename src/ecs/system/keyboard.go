package system

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
	log "github.com/sirupsen/logrus"
)

type KeyBoard struct {
	keyBindings Bindings
	keyStates   keyStates
	results     ListenResult
	reg         *registry.Registry
}

func NewKeyBoard(r *registry.Registry) *KeyBoard {
	res := &KeyBoard{keyBindings: NewDefaultBindings(), keyStates: newKeyStates(), results: NewListenResult(), reg: r}

	go func() {
		for {
			select {
			case pressed := <-res.results.FireCh:
				log.Tracef("fire key pressed, %v", pressed)
				res.reg.RLock()
				for _, e := range res.reg.Entities {
					e.Lock()
					if weapC := e.GetComponents(component.WeaponID); weapC != nil {
						if userWeap, ok := weapC[0].(*component.UserControlledWeapon); ok {
							userWeap.SetAutoAttack(pressed)
						}
					}
					e.Unlock()
				}
				res.reg.RUnlock()
			case changedVec := <-res.results.VectorCh:
				res.reg.RLock()
				for _, e := range res.reg.Entities {
					e.Lock()
					if e.HasTags(tag.User) {
						e = e.RemoveComponents(component.VectorID)
						if changedVec != nil {
							e = e.SetComponents(*changedVec)
						}
					}
					e.Unlock()
				}
				res.reg.RUnlock()
			case changedWeapDir := <-res.results.WeaponDirCh:
				log.Tracef("weap dir changed: %v", changedWeapDir)
				res.reg.RLock()
				for _, e := range res.reg.Entities {
					e.Lock()
					if weapC := e.GetComponents(component.WeaponID); weapC != nil {
						weap := weapC[0].(component.Weapon)
						if wdir, ok := weap.WeaponDir().(component.UserControlledWeaponDirection); ok {
							newDir := wdir.NewDirection(changedWeapDir)
							weap.SetDirection(newDir)
						}
					}
					e.Unlock()
				}
				res.reg.RUnlock()
			}
		}
	}()

	return res
}

func (k *KeyBoard) Step() {
	vectorChanged := false
	weaponDirChanged := false
	for eKey, Key := range k.keyBindings {
		eKey, Key := eKey, Key
		if inpututil.IsKeyJustPressed(eKey) {
			switch Key {
			case Fire:
				k.results.FireCh <- true
			case Up, Down, Left, Right:
				vectorChanged = true
				k.keyStates[Key] = true
			case WeapLeft, WeapRight:
				weaponDirChanged = true
				k.keyStates[Key] = true
			}
		} else if inpututil.IsKeyJustReleased(eKey) {
			switch Key {
			case Fire:
				k.results.FireCh <- false
			case Up, Down, Left, Right:
				vectorChanged = true
				k.keyStates[Key] = false
			case WeapLeft, WeapRight:
				weaponDirChanged = true
				k.keyStates[Key] = false
			}
		}
	}
	if vectorChanged {
		k.results.VectorCh <- getVector(k.keyStates[Up], k.keyStates[Down], k.keyStates[Left], k.keyStates[Right])
	}
	if weaponDirChanged {
		k.results.WeaponDirCh <- getWeaponDir(k.keyStates[WeapLeft], k.keyStates[WeapRight])
	}
}

// Key ...
type Key int

func (k Key) String() string {
	switch k {
	case Fire:
		return "Fire"
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return "Unknown key"
}

// keys
const (
	Fire Key = iota

	Up
	Down
	Left
	Right

	WeapLeft
	WeapRight
)

type ListenResult struct {
	FireCh      chan bool
	VectorCh    chan *component.Vector
	WeaponDirCh chan component.WeaponDir
}

func NewListenResult() ListenResult {
	return ListenResult{FireCh: make(chan bool), VectorCh: make(chan *component.Vector), WeaponDirCh: make(chan component.WeaponDir)}
}

// Bindings ...
type Bindings map[ebiten.Key]Key

// NewDefaultBindings ...
func NewDefaultBindings() Bindings {
	return Bindings(map[ebiten.Key]Key{
		ebiten.KeySpace: Fire,

		ebiten.KeyW: Up,
		ebiten.KeyS: Down,
		ebiten.KeyA: Left,
		ebiten.KeyD: Right,

		ebiten.KeyLeft:  WeapLeft,
		ebiten.KeyRight: WeapRight,
	})
}

// keyStates holds pressed/unpressed states of Keys
type keyStates map[Key]bool

func newKeyStates() keyStates {
	return keyStates(make(map[Key]bool))
}

func getWeaponDir(weapLeft, weapRight bool) component.WeaponDir {
	if weapLeft && !weapRight {
		return component.WeaponDirLeft
	}
	if !weapLeft && weapRight {
		return component.WeaponDirRight
	}
	return component.WeaponDirZero
}

// getVector is dumbest function ever
func getVector(up, down, left, right bool) *component.Vector {
	// up
	if left && up && right && !down || up && !left && !right && !down {
		res := component.Vector(1.5 * math.Pi)
		return &res
	}
	// down
	if left && down && right && !up || down && !left && !right && !up {
		res := component.Vector(0.5 * math.Pi)
		return &res
	}
	// left
	if left && up && down && !right || left && !up && !down && !right {
		res := component.Vector(math.Pi)
		return &res
	}
	// right
	if right && up && down && !left || right && !left && !up && !down {
		res := component.Vector(0)
		return &res
	}
	// left up
	if left && up && !right && !down {
		res := component.Vector(1.25 * math.Pi)
		return &res
	}
	// right up
	if right && up && !left && !down {
		res := component.Vector(1.75 * math.Pi)
		return &res
	}
	// right down
	if right && down && !left && !up {
		res := component.Vector(0.25 * math.Pi)
		return &res
	}
	// left down
	if left && down && !right && !up {
		res := component.Vector(0.75 * math.Pi)
		return &res
	}
	if up && down && left && right || !up && !down && !left && !right {
		return nil
	}
	if up && down && !left && !right || left && right && !up && !down {
		return nil
	}
	return nil
}
