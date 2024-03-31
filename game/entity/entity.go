package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Draw(image *ebiten.Image)
	Update()
}
