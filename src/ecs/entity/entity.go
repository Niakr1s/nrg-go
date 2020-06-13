package entity

import (
	"sync"

	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

type EntityID int

// Entity holds ID and Components
type Entity struct {
	ID EntityID
	sync.RWMutex

	Components map[component.ID]component.Component
	Tags       map[tag.ID]struct{}
}

// NewEntity constructs Entity with ID=0 and empty Components
func NewEntity() *Entity {
	return &Entity{Components: make(map[component.ID]component.Component), Tags: make(map[tag.ID]struct{})}
}

func (e *Entity) SetID(id EntityID) *Entity {
	e.ID = id
	return e
}

func (e *Entity) SetComponent(c component.Component) *Entity {
	e.Components[c.ID()] = c
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

func (e *Entity) SetTag(id tag.ID) *Entity {
	e.Tags[id] = struct{}{}
	return e
}

func (e *Entity) SetTags(ids ...tag.ID) *Entity {
	for _, id := range ids {
		e = e.SetTag(id)
	}
	return e
}

func (e *Entity) HasTag(id tag.ID) bool {
	_, ok := e.Tags[id]
	return ok
}

func (e *Entity) RemoveTag(id tag.ID) {
	delete(e.Tags, id)
}
