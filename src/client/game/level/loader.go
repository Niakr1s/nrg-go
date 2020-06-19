package level

import (
	"github.com/niakr1s/nrg-go/src/ecs/registry"
)

// InitLevelFunc is a function, that initializes a registry for level.
type InitLevelFunc func() *registry.Registry

type Loader struct {
	Reg *registry.Registry

	levelFuncs []InitLevelFunc
	current    int
}

func NewLoader() *Loader {
	return &Loader{levelFuncs: getInitLevelFuncs()}
}

func (l *Loader) LoadLevel() {
	l.Reg = l.levelFuncs[l.current]()
}

func (l *Loader) NextLevel() bool {
	l.current++
	if l.current >= len(l.levelFuncs) {
		return false
	}
	l.Reg = l.levelFuncs[l.current]()
	return true
}
