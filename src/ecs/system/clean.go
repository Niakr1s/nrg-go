package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

type Clean struct{}

func NewClean() *Clean {
	return &Clean{}
}

func (c *Clean) Step(reg *registry.Registry) {
	reg.Lock()
	defer reg.Unlock()

	setDestroyedTagsForDoneAnimations(reg.Entities)
	reg.Entities = removeDestroyed(reg.Entities)
}

func setDestroyedTagsForDoneAnimations(entities []*entity.Entity) {
	for _, e := range entities {
		e.Lock()
		if animC := e.GetComponents(component.AnimationID); animC != nil {
			anim := animC[0].(component.Animation)
			if anim.Done() {
				e.SetTags(tag.Destroyed)
			}
		}
		e.Unlock()
	}
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
