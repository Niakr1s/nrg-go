package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/img"
)

type Animation interface {
	Component

	// Step is used to notify Animation, that current frame was displayed.
	Step()
	GetImage() *ebiten.Image
	Done() bool
}

type BaseAnimation struct{}

func (a BaseAnimation) ID() ID {
	return AnimationID
}

type ExplodeAnimation struct {
	BaseAnimation

	maxFrames    int
	currentFrame int
	frames       []*ebiten.Image
}

func NewExplodeAnimation() *ExplodeAnimation {
	framesStr := []string{img.Explosion0, img.Explosion1, img.Explosion2}
	frames := make([]*ebiten.Image, len(framesStr))
	for i, frameStr := range framesStr {
		frame, err := img.Load(frameStr)
		if err != nil {
			panic(err)
		}
		frames[i] = frame
	}
	return &ExplodeAnimation{maxFrames: 30, frames: frames}
}

func (ea *ExplodeAnimation) Step() {
	ea.currentFrame++
}

func (ea *ExplodeAnimation) Done() bool {
	return ea.currentFrame >= ea.maxFrames
}

func (ea *ExplodeAnimation) GetImage() *ebiten.Image {
	if len(ea.frames) == 0 {
		return nil
	}
	framesPerImage := ea.maxFrames / len(ea.frames)
	frameNo := ea.currentFrame / framesPerImage
	// just in case
	if frameNo >= len(ea.frames) {
		frameNo = len(ea.frames) - 1
	}
	return ea.frames[frameNo]
}
