package main

import (
	cm "bob/libs/collision"
	em "bob/libs/entity"
	am "bob/libs/animation"
	. "bob/libs/types"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
	entity_list []*em.Entity
}

func (g *Game) Update() error {
	cm.ListCollision(collisionFromEntity(g.entity_list))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.entity_list{
		o.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 160
}

// create collision slice from entity slice
func collisionFromEntity(eSlice []*em.Entity) []*cm.CollisionBox{
	var t []*cm.CollisionBox
	for _, e := range eSlice{
		t = append(t, e.GetCollider())
	}
	return t
}

func main() {
	game := Game{}

	// create player entity

	player := em.Entity{}
	player_sprite := am.AnimatedImage{}
	player_sprite.Init("assets/player.png", 500, 16, 16)
	go player_sprite.PlayLoop()
	player_sprite.SetPlaying(true)
	
	player.Init(&player_sprite, &Vec2f{100.0, 100.0}, &Vec2f{0, 0}, true)

	player.MakeCollider(0, Vec2f{16, 16}, Vec2f{100, 100},
		func(id int) {
			log.Printf("Player collision with %v\n", id)
		},
	)

	player.FlipX(true)
	
	game.entity_list = append(game.entity_list, &player)

	// create enemy entity
	
	
	
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")
	
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
} 
