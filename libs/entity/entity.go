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

	flip_x bool
	flip_y bool
}

func (e *Entity) Init(drawble drawble.Drawble, position, velocity *Vec2f, has_collision bool) {
	e.drawble = drawble
	e.position = *position
	e.velocity = *velocity
	e.collider.Is_on = has_collision
}

func (e *Entity) MakeCollider(id int, size, position Vec2f, trigger cm.Trigger_function) {
	e.collider.Init(id, size, position, trigger)
}

func (e *Entity) GetCollider() *cm.CollisionBox {
	return &e.collider
}

func (e *Entity) Draw(surface *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	// flip x & y
	if e.flip_x {
		opt.GeoM.Scale(-1, 1)
		opt.GeoM.Translate(e.drawble.GetSize().X, 0)
	}
	opt.GeoM.Translate(e.position.X, e.position.Y)
	surface.DrawImage(e.drawble.GetImage(), opt)
}

func (e *Entity) FlipX(should_flip bool) {
	e.flip_x = should_flip
}

func (e *Entity) Update() {
	e.position.AddEql(&e.velocity)
	e.collider.SetPosition(e.position)
}

func (e *Entity) SetVelocity(vel Vec2f) {
	e.velocity = vel
}

func (e *Entity) GetVelocity() Vec2f { return e.velocity }

func (e *Entity) AddVelocity(add Vec2f) {
	e.velocity.AddEql(&add)
}

func (e *Entity) GetPosition() Vec2f { return e.position }

func (e *Entity) GetSize() Vec2f{
	return e.drawble.GetSize()
}
