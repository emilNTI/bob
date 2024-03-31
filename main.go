package main

import (
	em "bob/game/entity"
	"bob/game/player"
	am "bob/libs/animation"
	cm "bob/libs/collision"

	. "bob/libs/types"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	entity_list []em.Entity
	player *player.Player
}

func (g *Game) Update() error {
	for _, o := range g.entity_list{
		(o).Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{46, 34, 47, 255})
	for _, o := range g.entity_list {
		(o).Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 160
}

func main() {
	game := Game{}

	player_sprite := am.CreateAI("assets/player.png", 500, 16, 16)
	player_sprite.SetPlaying(true)
	p := player.CreatePlayer(
		player_sprite,
		cm.CreateCB(0, Vec2f{8, 3}, Vec2f{16, 16},
			func(id int) {
				log.Printf("Player collision with %v\n", id)
			},
		),
		cm.CreateCB(1, Vec2f{8, 16}, Vec2f{16, 16},
			func(id int) {
				log.Printf("Player box 2 collision with %v\n", id)
			},
		),
		Vec2f{90, 90},
	)
	
	game.player = &p

	game.entity_list = append(game.entity_list, &p)

		
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
