package system

import (
	"testing"

	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/stretchr/testify/assert"
)

func Test_correctPosBoard(t *testing.T) {
	const (
		boardW = 800
		boardH = 600
	)

	tests := []struct {
		name string

		pos  component.Pos
		w, h float64

		want component.Pos
	}{
		{"inside board", component.NewPos(500, 500), 100, 100, component.NewPos(500, 500)},

		{"partly left outside", component.NewPos(-25, 500), 100, 100, component.NewPos(50, 500)},
		{"partly right outside", component.NewPos(825, 500), 100, 100, component.NewPos(750, 500)},
		{"partly top outside", component.NewPos(500, -25), 100, 100, component.NewPos(500, 50)},
		{"partly bot outside", component.NewPos(500, 625), 100, 100, component.NewPos(500, 550)},
		{"partly left top outside", component.NewPos(-25, -25), 100, 100, component.NewPos(50, 50)},
		{"partly right top outside", component.NewPos(825, -25), 100, 100, component.NewPos(750, 50)},
		{"partly right bot outside", component.NewPos(825, 625), 100, 100, component.NewPos(750, 550)},
		{"partly left bot outside", component.NewPos(-25, 625), 100, 100, component.NewPos(50, 550)},

		{"full left outside", component.NewPos(-2000, 500), 100, 100, component.NewPos(50, 500)},
		{"full right outside", component.NewPos(2000, 500), 100, 100, component.NewPos(750, 500)},
		{"full top outside", component.NewPos(500, -2000), 100, 100, component.NewPos(500, 50)},
		{"full bot outside", component.NewPos(500, 2000), 100, 100, component.NewPos(500, 550)},
		{"full left top outside", component.NewPos(-2000, -2000), 100, 100, component.NewPos(50, 50)},
		{"full right top outside", component.NewPos(2000, -2000), 100, 100, component.NewPos(750, 50)},
		{"full right bot outside", component.NewPos(2000, 2000), 100, 100, component.NewPos(750, 550)},
		{"full left bot outside", component.NewPos(-2000, 2000), 100, 100, component.NewPos(50, 550)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bound := component.NewBound(tt.pos, tt.w, tt.h)
			got, _ := correctPosBoard(tt.pos, bound, boardW, boardH)
			assert.Equal(t, tt.want, got)
		})
	}
}
