package image

import (
	. "bob/libs/types"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Image struct {
	sprite *ebiten.Image
	size   Vec2f
}

func (i *Image) Init(path string, width, height uint) {
	i.size.X = float64(width)
	i.size.Y = float64(height)
	t, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal("Error loading in static image")
	}
	i.sprite = t
}

func (i *Image) GetImage() *ebiten.Image {
	return i.sprite
}

func (i *Image) GetSize() Vec2f {
	return i.size
}
