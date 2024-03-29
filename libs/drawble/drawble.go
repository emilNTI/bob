package drawble

import "github.com/hajimehoshi/ebiten/v2"

type Drawble interface {
	GetImage() *ebiten.Image
}
