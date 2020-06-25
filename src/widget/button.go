package widget

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type Button struct {
	w, h    int
	Text    string
	onClick func()
	f       font.Face
}

func NewButton(w, h int, text string, onClick func()) *Button {
	return &Button{w: w, h: h, Text: text, onClick: onClick, f: NewBigFont()}
}

func (b *Button) Width() int {
	return b.w
}

func (b *Button) Height() int {
	return b.h
}

func (b *Button) OnEvent(e Event) {
	if e.Type() == MouseClick {
		b.onClick()
	}
}

func (b *Button) Draw() *ebiten.Image {
	res, _ := ebiten.NewImage(b.w, b.h, ebiten.FilterDefault)
	res.Fill(color.White)
	x, y := b.getXYForCenteredText()
	text.Draw(res, b.Text, b.f, x, y, color.Black)
	return res
}

func (b *Button) getXYForCenteredText() (int, int) {
	w, h := b.getTextWidthAndHeight()
	x := (b.w - w) / 2
	y := (b.h-h)/2 + h
	y -= 5 // otherwise text doesn't centers vertically
	return x, y
}

func (b *Button) getTextWidthAndHeight() (int, int) {
	_, advance := font.BoundString(b.f, b.Text)
	w := advance.Round()
	h := b.f.Metrics().Ascent.Round()
	return w, h
}
