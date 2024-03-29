package main

import (
	entityManager "bob/libs"
	am "bob/libs/animation"
	. "bob/libs/types"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
	player entityManager.Entity
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 160
}

func main() {
	game := Game{}

	pE := am.AnimatedImage{}

	game.player.Init(&pE, &Vec2f{100.0, 100.0}, &Vec2f{0.0, 0.0})

	pE.Init("assets/player.png", 500, 16, 16)
	go pE.PlayLoop()

	pE.SetPlaying(true)
	
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")
	
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
} 
