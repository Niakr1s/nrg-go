package system

import (
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
)

type Status struct {
	LevelCompleted bool
	LevelFailed    bool
}

func NewStatus() *Status {
	return &Status{}
}

func (l *Status) Reset() {
	l.LevelCompleted = false
	l.LevelFailed = false
}

func (l *Status) Step(reg *registry.Registry) {
	reg.RLock()
	defer reg.RUnlock()

	enemyFound := false
	userFound := false
	for _, e := range reg.Entities {
		if e.HasTags(tag.User) {
			userFound = true
		}
		if e.HasTags(tag.Enemy) {
			enemyFound = true
		}
	}

	if !userFound {
		l.LevelFailed = true
	}
	if !enemyFound {
		l.LevelCompleted = true
	}
}
