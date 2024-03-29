package collision

import (
	. "bob/libs/types"
	"bob/libs/drawble"
	// "log"
)

// called when a object collides with something (id)
type trigger_function func(id int)

// position is top left corner
type CollisionBox struct {
	size     Vec2f
	position Vec2f
	trigger  trigger_function
	id       int // represents type
}

func (c *CollisionBox) Init(size, position Vec2f) {
	c.size = size
	c.position = position
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

func ListCollision(list []CollisionBox) {
	for i1, b1 := range list {
		for i2, b2 := range list {
			if i2 == i1 {
				continue
			}
			if b1.IsColliding(&b2) {
				b1.trigger(b2.id)
			}
		}
	}
}
