package key

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/niakr1s/nrg-go/src/geo"
)

// Key ...
type Key int

func (k Key) String() string {
	switch k {
	case Exit:
		return "Exit"
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
	Exit Key = iota

	Fire

	Up
	Down
	Left
	Right
)

type ListenResult struct {
	ExitCh   chan bool
	FireCh   chan bool
	VectorCh chan *geo.Vector
}

func NewListenResult() ListenResult {
	return ListenResult{ExitCh: make(chan bool), FireCh: make(chan bool), VectorCh: make(chan *geo.Vector)}
}

// Bindings ...
type Bindings map[ebiten.Key]Key

// NewDefaultBindings ...
func NewDefaultBindings() Bindings {
	return Bindings(map[ebiten.Key]Key{
		ebiten.KeyEscape: Exit,

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

// Listener ...
type Listener struct {
	keyBindings Bindings
	keyStates   keyStates

	out ListenResult
}

// NewListener ...
func NewListener() *Listener {
	return &Listener{keyBindings: NewDefaultBindings(), keyStates: newKeyStates(), out: NewListenResult()}
}

// StartPollKeys starts polling keys. frameCh is used to poll keyboard not faster than framerate.
func (l *Listener) StartPollKeys(frameCh <-chan struct{}) ListenResult {
	go l.loop(frameCh)
	return l.out
}

func (l *Listener) loop(frameCh <-chan struct{}) {
	for {
		<-frameCh
		vectorChanged := false
		for eKey, Key := range l.keyBindings {
			eKey, Key := eKey, Key
			if inpututil.IsKeyJustPressed(eKey) {
				switch Key {
				case Fire:
					l.out.FireCh <- true
				case Exit:
					l.out.ExitCh <- true
				case Up, Down, Left, Right:
					vectorChanged = true
					l.keyStates[Key] = true
				}
			} else if inpututil.IsKeyJustReleased(eKey) {
				switch Key {
				case Fire:
					l.out.FireCh <- false
				case Exit:
					l.out.ExitCh <- false
				case Up, Down, Left, Right:
					vectorChanged = true
					l.keyStates[Key] = false
				}
			}
		}
		if vectorChanged {
			l.out.VectorCh <- getVector(l.keyStates[Up], l.keyStates[Down], l.keyStates[Left], l.keyStates[Right])
		}
	}
}

// getVector is dumbest function ever
func getVector(up, down, left, right bool) *geo.Vector {
	// up
	if left && up && right && !down || up && !left && !right && !down {
		res := geo.Vector(1.5 * math.Pi)
		return &res
	}
	// down
	if left && down && right && !up || down && !left && !right && !up {
		res := geo.Vector(0.5 * math.Pi)
		return &res
	}
	// left
	if left && up && down && !right || left && !up && !down && !right {
		res := geo.Vector(math.Pi)
		return &res
	}
	// right
	if right && up && down && !left || right && !left && !up && !down {
		res := geo.Vector(0)
		return &res
	}
	// left up
	if left && up && !right && !down {
		res := geo.Vector(1.25 * math.Pi)
		return &res
	}
	// right up
	if right && up && !left && !down {
		res := geo.Vector(1.75 * math.Pi)
		return &res
	}
	// right down
	if right && down && !left && !up {
		res := geo.Vector(0.25 * math.Pi)
		return &res
	}
	// left down
	if left && down && !right && !up {
		res := geo.Vector(0.75 * math.Pi)
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
