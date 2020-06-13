package system

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVector(t *testing.T) {
	type args struct {
		right bool
		down  bool
		left  bool
		up    bool
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantRad float64
	}{
		{"none pressed", args{false, false, false, false}, true, 0},
		{"all pressed", args{true, true, true, true}, true, 0},
		{"updown", args{false, true, false, true}, true, 0},
		{"leftright", args{true, false, true, false}, true, 0},
		{"right1", args{true, false, false, false}, false, 0},
		{"right2", args{true, true, false, true}, false, 0},
		{"down1", args{false, true, false, false}, false, 0.5 * math.Pi},
		{"down2", args{true, true, true, false}, false, 0.5 * math.Pi},
		{"left1", args{false, false, true, false}, false, math.Pi},
		{"left2", args{false, true, true, true}, false, math.Pi},
		{"up1", args{false, false, false, true}, false, 1.5 * math.Pi},
		{"up2", args{true, false, true, true}, false, 1.5 * math.Pi},
		{"rightdown", args{true, true, false, false}, false, 0.25 * math.Pi},
		{"leftdown", args{false, true, true, false}, false, 0.75 * math.Pi},
		{"leftup", args{false, false, true, true}, false, 1.25 * math.Pi},
		{"rightup", args{true, false, false, true}, false, 1.75 * math.Pi},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getVector(tt.args.up, tt.args.down, tt.args.left, tt.args.right)
			if tt.wantNil {
				assert.Nil(t, got)
				return
			}
			assert.InDelta(t, tt.wantRad, float64(*got), 0.1)
		})
	}
}
