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


type Vec2i struct{
	X int
	Y int
}

func (v1 *Vec2i) AddEql(v2 *Vec2i){
	v1.X += v2.X
	v1.Y += v2.Y
}
