package entityManager

import (
	cm "bob/libs/collision"
	"bob/libs/drawble"
	. "bob/libs/types"
	//"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	drawble  drawble.Drawble
	position Vec2f
	velocity Vec2f
	collider cm.CollisionBox
}

func (e *Entity) Init(drawble drawble.Drawble, position, velocity *Vec2f, has_collision bool) {
	e.drawble = drawble
	e.position = *position
	e.velocity = *velocity
	e.collider.Is_on = has_collision
}

func (e *Entity) MakeCollider(id int, size, position *Vec2f, trigger cm.Trigger_function){
	e.collider.Init(id, *size, *position, trigger)
}

func (e *Entity) Draw(surface *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(e.position.X, e.position.Y)
	surface.DrawImage(e.drawble.GetImage(), opt)
}

func (e *Entity) Update() {
	e.position.AddEql(&e.velocity)
	e.collider.SetPosition(e.position)
}
