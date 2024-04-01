package main

import (
	"bob/game/enemy"
	em "bob/game/entity"
	"bob/game/player"
	am "bob/libs/animation"
	cm "bob/libs/collision"
	"fmt"
	"math/rand"

	. "bob/libs/types"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	entity_list []em.Entity
	player *player.Player
	col []*cm.CollisionBox
}

func (g *Game) Update() error {
	for _, o := range g.entity_list{
		(o).Update()
	}
	cm.ListCollision(g.col)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{46, 34, 47, 255})
	
	for _, c := range g.col {
		c.Draw(screen, color.RGBA{255, uint8(rand.Intn(255)), 0, 255})
	}
	
	for _, o := range g.entity_list {
		(o).Draw(screen)
	}
	
	fps := fmt.Sprintf("FPS: %v", ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, fps)
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

	enemy_sprite := am.CreateAI("assets/enemy.png", 500, 16, 16)
	enemy_sprite.SetPlaying(true)
	e := enemy.CreateSkeleton(enemy_sprite, Vec2f{100, 100}, &p)

	game.entity_list = append(game.entity_list, &e)

	game.col = append(game.col, p.GetHitBox())
	game.col = append(game.col, e.GetHitBox())
	
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
