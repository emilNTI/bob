package drawble

import (
	. "bob/libs/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawble interface {
	GetImage() *ebiten.Image
	GetSize() Vec2f
}
