package system

import (
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
	"github.com/stretchr/testify/assert"
)

func Test_removeDestroyed(t *testing.T) {
	type args struct {
		entities []*entity.Entity
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{"empty", args{[]*entity.Entity{}}, 0},
		{"one not destroyed", args{[]*entity.Entity{
			entity.NewEntity(),
		}}, 1},
		{"one destroyed", args{[]*entity.Entity{
			entity.NewEntity().SetTags(tag.Destroyed),
		}}, 0},
		{"two not destroyed", args{[]*entity.Entity{
			entity.NewEntity(),
			entity.NewEntity(),
		}}, 2},
		{"two destroyed", args{[]*entity.Entity{
			entity.NewEntity().SetTags(tag.Destroyed),
			entity.NewEntity().SetTags(tag.Destroyed),
		}}, 0},
		{"two: destroyed and not destroyed", args{[]*entity.Entity{
			entity.NewEntity(),
			entity.NewEntity().SetTags(tag.Destroyed),
		}}, 1},
		{"random case", args{[]*entity.Entity{
			entity.NewEntity(),
			entity.NewEntity().SetTags(tag.Destroyed),
			entity.NewEntity(),
			entity.NewEntity().SetTags(tag.Destroyed),
			entity.NewEntity(),
			entity.NewEntity().SetTags(tag.Destroyed),
			entity.NewEntity(),
			entity.NewEntity().SetTags(tag.Destroyed),
			entity.NewEntity().SetTags(tag.Destroyed),
		}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDestroyed(tt.args.entities)
			assert.Len(t, got, tt.wantLen)
		})
	}
}

type mockAnimation struct {
	component.BaseAnimation
	done bool
}

func (a *mockAnimation) Step()                   {}
func (a *mockAnimation) GetImage() *ebiten.Image { return nil }
func (a *mockAnimation) Done() bool              { return a.done }

func Test_addDestroyedTagsForDoneAnimations(t *testing.T) {
	type args struct {
		anim component.Animation
	}
	tests := []struct {
		name          string
		args          args
		wantDestroyed bool
	}{
		{"done animation", args{anim: &mockAnimation{done: true}}, true},
		{"not done animation", args{anim: &mockAnimation{done: false}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := registry.NewRegistry()
			reg.AddEntity(entity.NewEntity().SetComponents(tt.args.anim))
			setDestroyedTagsForDoneAnimations(reg.Entities)
			assert.Equal(t, tt.wantDestroyed, reg.Entities[0].HasTags(tag.Destroyed))
		})
	}
}
