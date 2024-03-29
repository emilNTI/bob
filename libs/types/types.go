package types

// Vec2d float type
type Vec2f struct {
	X float64
	Y float64
}

func (v1 *Vec2f) AddEql(v2 *Vec2f){
	v1.X += v2.X
	v1.Y += v2.Y
}
