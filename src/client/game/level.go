package game

import (
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

// Level implements 2 methods of Level interface: GetRegistry.
// Inheritors must implement others.
type Level struct {
	Reg *registry.Registry
}

func NewLevel() *Level {
	return &Level{}
}

func (b *Level) GetRegistry() *registry.Registry {
	return b.Reg
}
