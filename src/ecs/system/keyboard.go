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
			case changedVec := <-res.results.VectorCh:
				res.reg.RLock()
				for _, e := range res.reg.Entities {
					e.RLock()
					if e.HasTags(tag.UserID) {
						e.RUnlock()
						e.Lock()
						e = e.RemoveComponents(component.VectorID)
						if changedVec != nil {
							e = e.SetComponents(changedVec)
						}
						e.Unlock()
						e.RLock()
					}
					e.RUnlock()
				}
				res.reg.RUnlock()
			}
		}
	}()

	return res
}

func (k *KeyBoard) Step() {
	vectorChanged := false
	for eKey, Key := range k.keyBindings {
		eKey, Key := eKey, Key
		if inpututil.IsKeyJustPressed(eKey) {
			switch Key {
			case Fire:
				k.results.FireCh <- true
			case Up, Down, Left, Right:
				vectorChanged = true
				k.keyStates[Key] = true
			}
		} else if inpututil.IsKeyJustReleased(eKey) {
			switch Key {
			case Fire:
				k.results.FireCh <- false
			case Up, Down, Left, Right:
				vectorChanged = true
				k.keyStates[Key] = false
			}
		}
	}
	if vectorChanged {
		k.results.VectorCh <- getVector(k.keyStates[Up], k.keyStates[Down], k.keyStates[Left], k.keyStates[Right])
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
)

type ListenResult struct {
	FireCh   chan bool
	VectorCh chan *component.Vector
}

func NewListenResult() ListenResult {
	return ListenResult{FireCh: make(chan bool), VectorCh: make(chan *component.Vector)}
}

// Bindings ...
type Bindings map[ebiten.Key]Key

// NewDefaultBindings ...
func NewDefaultBindings() Bindings {
	return Bindings(map[ebiten.Key]Key{
		ebiten.KeySpace: Fire,

		ebiten.KeyUp: Up,
		ebiten.KeyW:  Up,

		ebiten.KeyDown: Down,
		ebiten.KeyS:    Down,

		ebiten.KeyLeft: Left,
		ebiten.KeyA:    Left,

		ebiten.KeyRight: Right,
		ebiten.KeyD:     Right,
	})
}

// keyStates holds pressed/unpressed states of Keys
type keyStates map[Key]bool

func newKeyStates() keyStates {
	return keyStates(make(map[Key]bool))
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
