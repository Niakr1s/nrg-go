package component

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle_OuterPointInDirection(t *testing.T) {

	type args struct {
		circle Circle
		center Pos
		vec    Vector
	}
	tests := []struct {
		name string
		args args
		want Pos
	}{
		{"right", args{NewCircle(10), NewPos(0, 0), NewVector(0)}, NewPos(10, 0)},
		{"right2", args{NewCircle(10), NewPos(10, 10), NewVector(0)}, NewPos(20, 10)},
		{"left", args{NewCircle(10), NewPos(0, 0), NewVector(math.Pi)}, NewPos(-10, 0)},
		{"left2", args{NewCircle(10), NewPos(10, 10), NewVector((math.Pi))}, NewPos(0, 10)},
		{"top", args{NewCircle(10), NewPos(0, 0), NewVector(1.5 * math.Pi)}, NewPos(0, -10)},
		{"top2", args{NewCircle(10), NewPos(10, 10), NewVector(1.5 * math.Pi)}, NewPos(10, 0)},
		{"down", args{NewCircle(10), NewPos(0, 0), NewVector(0.5 * math.Pi)}, NewPos(0, 10)},
		{"down2", args{NewCircle(10), NewPos(10, 10), NewVector(0.5 * math.Pi)}, NewPos(10, 20)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.circle.OuterPointInDirection(tt.args.center, tt.args.vec)
			assert.InDelta(t, tt.want.X, got.X, 0.01)
			assert.InDelta(t, tt.want.Y, got.Y, 0.01)
		})
	}
}
