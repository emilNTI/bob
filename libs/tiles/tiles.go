package tileManager

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	img ebiten.Image
	id  uint
}

type TileManager struct {
	draw_on    *ebiten.Image
	tiles      []Tile
	game_map   []uint
	game_map_w uint
	game_map_h uint
}

func (t *TileManager) Init(draw_on *ebiten.Image) {
	t.draw_on = draw_on
}

func (t *TileManager) AddTile(path string){
	temp, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil{
		log.Fatal("Failed to load tile")
	}
	t.tiles = append(t.tiles, Tile{*temp, 0})
}

func (t *TileManager) getTileAt(x, y int) (*Tile){
	return &t.tiles[y * int(t.game_map_w) + x]
}
