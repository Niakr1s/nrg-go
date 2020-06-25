package widget

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

var defaultFont *truetype.Font

const dpi = 72

func init() {
	var err error
	defaultFont, err = truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
}

func NewNormalFont() font.Face {
	return NewFont(24)
}

func NewBigFont() font.Face {
	return NewFont(48)
}

func NewFont(size float64) font.Face {
	return truetype.NewFace(defaultFont, &truetype.Options{
		DPI:     dpi,
		Size:    size,
		Hinting: font.HintingFull,
	})
}
