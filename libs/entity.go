package entityManager

import (
	"bob/libs/drawble"
	. "bob/libs/types"
	//"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	drawble          drawble.Drawble
	position         Vec2f
	velocity         Vec2f
}

func (e *Entity) Init(drawble drawble.Drawble, position, velocity *Vec2f){
	e.drawble = drawble
	e.position = *position
	e.velocity = *velocity
}

func (e *Entity) Draw(surface *ebiten.Image){
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(e.position.X, e.position.Y)
	surface.DrawImage(e.drawble.GetImage(), opt)
}

func (e *Entity) Update(){
	e.position.AddEql(&e.velocity)
}
