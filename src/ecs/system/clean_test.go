package system

import (
	"testing"

	"github.com/niakr1s/nrg-go/src/ecs/entity"
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
