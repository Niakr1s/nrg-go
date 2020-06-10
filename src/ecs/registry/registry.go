package registry

import "github.com/niakr1s/nrg-go/src/ecs/entity"

// Registry contains entities
type Registry struct {
	lastID entity.EntityID

	Entities map[entity.EntityID]*entity.Entity
}

// NewRegistry constructs Registry
func NewRegistry() *Registry {
	return &Registry{Entities: make(map[entity.EntityID]*entity.Entity)}
}

// AddEntity adds entity and assigns it new ID
func (r *Registry) AddEntity(e *entity.Entity) {
	e.ID = r.lastID
	r.Entities[r.lastID] = e
	r.lastID++
}

// SetEntity sets/updates entity using entities' ID
func (r *Registry) SetEntity(e *entity.Entity) {
	r.Entities[e.ID] = e
}

// GetEntity gets entity, can return nil
func (r *Registry) GetEntity(id entity.EntityID) *entity.Entity {
	return r.Entities[id]
}

// RemoveEntity just removes entity
func (r *Registry) RemoveEntity(id entity.EntityID) {
	delete(r.Entities, id)
}
