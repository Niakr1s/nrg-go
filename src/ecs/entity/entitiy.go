package entity

import "github.com/niakr1s/nrg-go/src/ecs/component"

type ComponentID int
type EntityID int

const (
	Drawable ComponentID = iota
	Intersectable
	Movable
)

// Entity holds ID and Components
type Entity struct {
	ID EntityID

	Components map[ComponentID]component.Component
}

// NewEntity constructs Entity with ID=0 and empty Components
func NewEntity() *Entity {
	return &Entity{Components: make(map[ComponentID]component.Component)}
}

func (e *Entity) WithID(id EntityID) *Entity {
	e.ID = id
	return e
}

func (e *Entity) WithComponent(id ComponentID, c component.Component) *Entity {
	e.Components[id] = c
	return e
}

func (e *Entity) RemoveComponent(id ComponentID) *Entity {
	delete(e.Components, id)
	return e
}
