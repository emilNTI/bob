package collisionManager

import (
	. "bob/libs/types"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// called when a object collides with something (id)
type Trigger_function func(id int)

// position is top left corner
type CollisionBox struct {
	size     Vec2f
	position Vec2f
	trigger  Trigger_function // public right now until better way is found
	id       int              // represents type
	Is_on    bool
}

func (c *CollisionBox) Init(id int, size, position Vec2f, trigger Trigger_function) {
	c.id = id
	c.size = size
	c.position = position
	c.trigger = trigger
}

func CreateCB(id int, size, position Vec2f, trigger Trigger_function) CollisionBox {
	c := CollisionBox{}
	c.Init(id, size, position, trigger)
	return c
}

func (c1 *CollisionBox) IsColliding(c2 *CollisionBox) bool {
	// AABB collision
	if c1.position.X < c2.position.X+c2.size.X &&
		c1.position.X+c1.size.X > c2.position.X &&
		c1.position.Y < c2.position.Y+c2.size.Y &&
		c1.position.Y+c1.size.Y > c2.position.Y {
		return true
	}
	return false
}

func (c *CollisionBox) SetPosition(new_pos Vec2f) {
	c.position = new_pos
}

func (c *CollisionBox) SetSize(new_size Vec2f) {
	c.size = new_size
}

// only for debug
func (c *CollisionBox) Draw(img *ebiten.Image, col color.RGBA) {
	t := ebiten.NewImage(int(c.size.X), int(c.size.Y))
	t.Fill(col)
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(c.position.X, c.position.Y)

	img.DrawImage(t, &opt)
}

// skip used to skip index in list
func (c *CollisionBox) CheckWithList(list []CollisionBox, skip int){
	for i, b := range list{
		if i == skip{continue}
		if c.IsColliding(&b){
			c.trigger(b.id)
		}
	}
}

func ListCollision(list []*CollisionBox) {
	for i1, b1 := range list {
		for i2, b2 := range list {
			if i2 == i1 {
				continue
			}
			if b1.IsColliding(b2) {
				b1.trigger(b2.id)
			}
		}
	}
}
