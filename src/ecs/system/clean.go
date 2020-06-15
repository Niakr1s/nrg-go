package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

type Clean struct {
	reg *registry.Registry
}

func NewClean(reg *registry.Registry) *Clean {
	return &Clean{reg: reg}
}

func (c *Clean) Step() {
	c.reg.Lock()
	defer c.reg.Unlock()

	c.reg.Entities = removeDestroyed(c.reg.Entities)
}

func removeDestroyed(entities []*entity.Entity) []*entity.Entity {
	l := len(entities)
	for i := 0; i < l; i++ {
		e := entities[i]
		e.RLock()
		if !e.HasTags(tag.Destroyed) {
			e.RUnlock()
			continue
		}
		// swapping to the end
		if e.HasTags(tag.Destroyed) {
			entities[l-1], entities[i] = entities[i], entities[l-1]
			l--
			i--
		}

		e.RUnlock()
	}
	return entities[:l]
}
