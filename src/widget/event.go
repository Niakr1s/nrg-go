package widget

import "github.com/niakr1s/nrg-go/src/ecs/component"

type EventType int

const (
	MouseOver EventType = iota
	MouseClick
)

type Event interface {
	Type() EventType
}

type PosEvent struct {
	t   EventType
	pos component.Pos
}

func NewMouseClickEvent(pos component.Pos) PosEvent {
	return PosEvent{t: MouseClick, pos: pos}
}

func NewMouseOverEvent(pos component.Pos) PosEvent {
	return PosEvent{t: MouseOver, pos: pos}
}

func (pe PosEvent) Type() EventType {
	return pe.t
}

func (pe PosEvent) Pos() component.Pos {
	return pe.pos
}
