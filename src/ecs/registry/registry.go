package registry

import (
	"sync"

	"github.com/niakr1s/nrg-go/src/ecs/entity"
)

// Registry contains entities
type Registry struct {
	sync.RWMutex
	lastID entity.EntityID

	Entities []*entity.Entity
}

// NewRegistry constructs Registry
func NewRegistry() *Registry {
	return &Registry{Entities: []*entity.Entity{}}
}

// AddEntity adds entity and assigns it new ID
func (r *Registry) AddEntity(e *entity.Entity) *Registry {
	e.ID = r.lastID
	r.Entities = append(r.Entities, e)
	r.lastID++
	return r
}
