package component

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle_OuterPointInDirection(t *testing.T) {

	type args struct {
		circle Circle
		vec    Vector
	}
	tests := []struct {
		name string
		args args
		want Pos
	}{
		{"right", args{NewCircle(10), NewVector(0)}, NewPos(10, 0)},
		{"left", args{NewCircle(10), NewVector(math.Pi)}, NewPos(-10, 0)},
		{"top", args{NewCircle(10), NewVector(1.5 * math.Pi)}, NewPos(0, -10)},
		{"down", args{NewCircle(10), NewVector(0.5 * math.Pi)}, NewPos(0, 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.circle.OuterPointInDirectionDiff(tt.args.vec)
			assert.InDelta(t, tt.want.X, got.X, 0.01)
			assert.InDelta(t, tt.want.Y, got.Y, 0.01)
		})
	}
}
