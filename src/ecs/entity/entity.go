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
	Tags       tag.Tags
}

// NewEntity constructs Entity with ID=0 and empty Components
func NewEntity() *Entity {
	return &Entity{Components: make(map[component.ID]component.Component), Tags: make(tag.Tags)}
}

func (e *Entity) SetID(id EntityID) *Entity {
	e.ID = id
	return e
}

func (e *Entity) SetComponents(cs ...component.Component) *Entity {
	for _, c := range cs {
		e.Components[c.ID()] = c
	}
	return e
}

// GetComponents gets all components, return nil if any is missing
// garantees len(result) == len(ids) on succes
func (e *Entity) GetComponents(ids ...component.ID) component.Components {
	res := component.Components{}
	for _, id := range ids {
		c, ok := e.Components[id]
		if !ok {
			return nil
		}
		res = append(res, c)
	}
	return res
}

func (e *Entity) RemoveComponents(ids ...component.ID) *Entity {
	for _, id := range ids {
		delete(e.Components, id)
	}
	return e
}

func (e *Entity) SetTags(ids ...tag.ID) *Entity {
	e.Tags.SetTags(ids...)
	return e
}

// HasTags returns false if any is missing
func (e *Entity) HasTags(ids ...tag.ID) bool {
	return e.Tags.HasTags(ids...)
}

func (e *Entity) RemoveTags(ids ...tag.ID) {
	e.Tags.RemoveTags(ids...)
}

func (e *Entity) CopyTags() tag.Tags {
	return e.Tags.CopyTags()
}
