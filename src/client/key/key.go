package key

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

// Event ...
type Event struct {
	Key     Key
	Pressed bool
}

func (e Event) String() string {
	if e.Pressed {
		return fmt.Sprintf("key %s pressed", e.Key)
	}
	return fmt.Sprintf("key %s released", e.Key)
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

	out chan Event
}

// NewListener ...
func NewListener() *Listener {
	return &Listener{keyBindings: NewDefaultBindings(), keyStates: newKeyStates(), out: make(chan Event, 10)}
}

// StartPollKeys ...
func (l *Listener) StartPollKeys() <-chan Event {
	go l.loop()
	return l.out
}

func (l *Listener) loop() {
	for {
		for eKey, Key := range l.keyBindings {
			if inpututil.IsKeyJustPressed(eKey) && !l.keyStates[Key] {
				l.keyStates[Key] = true
				l.out <- Event{Key, true}
			} else if inpututil.IsKeyJustReleased(eKey) && l.keyStates[Key] {
				l.keyStates[Key] = false
				l.out <- Event{Key, false}
			}
		}
	}
}
