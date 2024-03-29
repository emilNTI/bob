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
	entity_list []em.Entity
}

func (g *Game) Update() error {
	
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

func main() {
	game := Game{}

	// create player entity
	
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")
	
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
} 
