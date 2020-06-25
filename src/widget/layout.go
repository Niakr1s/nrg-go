package widget

import (
	"github.com/niakr1s/nrg-go/src/ecs/component"
)

type Layout []WidgetWithPos

type WidgetWithPos struct {
	Widget
	Pos component.Pos
}

func NewLayout() Layout {
	return make(Layout, 0)
}

func (l *Layout) AddWidget(pos component.Pos, w Widget) {
	wp := WidgetWithPos{w, pos}
	*l = append(*l, wp)
}

func (l Layout) NewEvent(e Event) {
	for _, w := range l {
		switch e := e.(type) {
		case PosEvent:
			if isInside(e.pos, w.Pos, w.Width(), w.Height()) {
				w.OnEvent(e)
			}
		}
	}
}

func isInside(point component.Pos, rectTopLeft component.Pos, w, h int) bool {
	return point.X >= rectTopLeft.X && point.X <= (rectTopLeft.X+float64(w)) && point.Y >= rectTopLeft.Y && point.Y <= rectTopLeft.Y+float64(h)
}
