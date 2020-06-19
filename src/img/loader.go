package img

import (
	"image"

	"os"

	// init png
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

var cache map[string]*ebiten.Image

// paths
const (
	WhiteCircle = "data/white_circle.png"
	RedCircle   = "data/red_circle.png"
	BlueCircle  = "data/blue_circle.png"

	Explosion0 = "data/explosion/explosion0.png"
	Explosion1 = "data/explosion/explosion1.png"
	Explosion2 = "data/explosion/explosion2.png"
)

func init() {
	cache = make(map[string]*ebiten.Image)
}

func Load(imagePath string) (*ebiten.Image, error) {
	if i, ok := cache[imagePath]; ok {
		return i, nil
	}
	reader, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	eImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	cache[imagePath] = eImg
	return eImg, nil
}
