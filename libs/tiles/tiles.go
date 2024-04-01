package tileManager

import (
	. "bob/libs/types"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	sub_image Vec2i
	id        int
	name      string
}

type TileManager struct{
	tiles []Tile
	draw_on *ebiten.Image
	game_map []int // array of ints representing tiles
	map_size Vec2i
}

func CreateTileManager(map_size Vec2i, draw_on *ebiten.Image) TileManager
