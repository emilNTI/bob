package collision

import (
	. "bob/libs/types"
	// "log"
)

type trigger_function func()

// position is top left corner
type CollisionBox struct{
	size Vec2f
	position Vec2f
	trigger trigger_function
}

func (c *CollisionBox) Init(size, position Vec2f){
	c.size = size
	c.position = position
}
