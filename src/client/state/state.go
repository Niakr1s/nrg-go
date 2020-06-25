package state

import "github.com/hajimehoshi/ebiten"

type State interface {
	Update(screen *ebiten.Image) error
	Draw(screen *ebiten.Image)
	Next
}

type Next interface {
	Next() State
	ClearNext()
}

type NextState struct {
	next State
}

func NewNext() *NextState {
	return &NextState{}
}

func (n *NextState) Next() State {
	return n.next
}

func (n *NextState) ClearNext() {
	n.next = nil
}

func (n *NextState) SetNext(s State) {
	n.next = s
}
