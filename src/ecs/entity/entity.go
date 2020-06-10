package entity

import "github.com/niakr1s/nrg-go/src/ecs/component"

type EntityID int

// Entity holds ID and Components
type Entity struct {
	ID EntityID

	Components map[component.ID]component.Component
}

// NewEntity constructs Entity with ID=0 and empty Components
func NewEntity() *Entity {
	return &Entity{Components: make(map[component.ID]component.Component)}
}

func (e *Entity) WithID(id EntityID) *Entity {
	e.ID = id
	return e
}

func (e *Entity) WithComponent(id component.ID, c component.Component) *Entity {
	e.Components[id] = c
	return e
}

// GetComponent gets component, can return nil
func (e *Entity) GetComponent(id component.ID) component.Component {
	if c, ok := e.Components[id]; ok {
		return c
	}
	return nil
}

func (e *Entity) RemoveComponent(id component.ID) *Entity {
	delete(e.Components, id)
	return e
}

func (e *Entity) WithTag(id component.ID) *Entity {
	e.Components[id] = nil
	return e
}

func (e *Entity) WithTags(ids ...component.ID) *Entity {
	for _, id := range ids {
		e = e.WithTag(id)
	}
	return e
}

func (e *Entity) HasTag(id component.ID) bool {
	_, ok := e.Components[id]
	return ok
}

func (e *Entity) RemoveTag(id component.ID) {
	delete(e.Components, id)
}
