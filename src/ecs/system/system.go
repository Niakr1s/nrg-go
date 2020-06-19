package system

import "github.com/niakr1s/nrg-go/src/ecs/registry"

type System interface {
	Step(reg *registry.Registry)
}
