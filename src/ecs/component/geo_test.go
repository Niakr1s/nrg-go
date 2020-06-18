package component

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVectorFromPos(t *testing.T) {
	type args struct {
		pos1 Pos
		pos2 Pos
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"right", args{NewPos(0, 0), NewPos(10, 0)}, NewVector(0)},
		{"left", args{NewPos(0, 0), NewPos(-10, 0)}, NewVector(math.Pi)},
		{"top", args{NewPos(0, 0), NewPos(0, -10)}, NewVector(math.Pi * (-0.5))},
		{"bot", args{NewPos(0, 0), NewPos(0, 10)}, NewVector(math.Pi * 0.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVectorFromPos(tt.args.pos1, tt.args.pos2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVector_Mirrored(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want Vector
	}{
		{"from right", NewVector(0), NewVector(math.Pi)},
		{"from left", NewVector(math.Pi), NewVector(0)},
		{"from top", NewVector(1.5 * math.Pi), NewVector(0.5 * math.Pi)},
		{"from bot", NewVector(0.5 * math.Pi), NewVector(1.5 * math.Pi)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Mirrored()
			assert.InDelta(t, tt.want.Cos(), got.Cos(), 0.01)
			assert.InDelta(t, tt.want.Sin(), got.Sin(), 0.01)
		})
	}
}
