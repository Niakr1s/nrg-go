package entity

import "github.com/niakr1s/nrg-go/src/ecs/component"

type Parent struct {
	Parent EntityID
}

func NewParent(parent EntityID) Parent {
	return Parent{Parent: parent}
}

func (p Parent) ID() component.ID {
	return component.ParentID
}
