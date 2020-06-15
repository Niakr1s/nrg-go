package system

import (
	"testing"

	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
	"github.com/stretchr/testify/assert"
)

func TestDamage_Step(t *testing.T) {
	tests := []struct {
		name           string
		bullet         *entity.Entity
		user           *entity.Entity
		wantHpDecrease int
	}{
		{"should deal damage #1",
			entity.NewEntity().SetComponents(component.NewPos(500, 500),
				&component.Circle{R: 1}, component.NewDamage(10)).
				SetTags(tag.Bullet),
			entity.NewEntity().SetComponents(component.NewPos(500, 500),
				&component.Circle{R: 10}, component.NewHP(100)).
				SetTags(tag.User),
			10,
		},
		{"should deal damage #2",
			entity.NewEntity().SetComponents(component.NewPos(500, 505),
				&component.Circle{R: 1}, component.NewDamage(10)).
				SetTags(tag.Bullet),
			entity.NewEntity().SetComponents(component.NewPos(500, 500),
				&component.Circle{R: 10}, component.NewHP(100)).
				SetTags(tag.User),
			10,
		},
		{"should deal damage #3",
			entity.NewEntity().SetComponents(component.NewPos(500, 510),
				&component.Circle{R: 1}, component.NewDamage(10)).
				SetTags(tag.Bullet),
			entity.NewEntity().SetComponents(component.NewPos(500, 500),
				&component.Circle{R: 10}, component.NewHP(100)).
				SetTags(tag.User),
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTest := func(r *registry.Registry) {
				startHp := tt.user.Components[component.HpID].(component.HP).Current
				d := NewDamage(r)
				d.Step()
				endHp := tt.user.Components[component.HpID].(component.HP).Current
				assert.Equal(t, tt.wantHpDecrease, startHp-endHp)
			}
			r := registry.NewRegistry().AddEntity(tt.bullet).AddEntity(tt.user)
			runTest(r)
			r.Entities[0], r.Entities[1] = r.Entities[1], r.Entities[0]
			runTest(r)
		})
	}
}
