package img

import (
	"image"

	"os"

	// init png
	_ "image/png"
)

var cache map[string]image.Image

// paths
const (
	WhiteCircle = "data/white_circle.png"
	RedCircle   = "data/red_circle.png"
	BlueCircle  = "data/blue_circle.png"
)

func init() {
	cache = make(map[string]image.Image)
}

func Load(imagePath string) (image.Image, error) {
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
	cache[imagePath] = img
	return img, nil
}
