package main

import (
	am "bob/libs/animation"
	cm "bob/libs/collision"
	em "bob/libs/entity"
	"bob/libs/image"
	"math"

	// im "bob/libs/image"
	. "bob/libs/types"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	entity_list    []*em.Entity
	player         em.Entity
	shoot_cooldown int
}

func (g *Game) Update() error {
	cm.ListCollision(collisionFromEntity(g.entity_list))

	// input
	if math.Abs(g.player.GetVelocity().Y) < 1.5 {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			g.player.AddVelocity(Vec2f{0, -0.2})
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.player.AddVelocity(Vec2f{0, 0.2})
		}
	}

	if math.Abs(g.player.GetVelocity().X) < 1.5 {
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			g.player.AddVelocity(Vec2f{0.2, 0})
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.player.AddVelocity(Vec2f{-0.2, 0})
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.shoot_cooldown <= 0{
		g.shoot_cooldown = 60
		createBullet(g)
	}

	for _, e := range g.entity_list {
		e.Update()
	}
	// mouse
	mx, _ := ebiten.CursorPosition()
	if mx > int(g.player.GetPosition().X+g.player.GetSize().X/2) {
		g.player.FlipX(false)
	} else {
		g.player.FlipX(true)
	}
	// break player

	if math.Abs(g.player.GetVelocity().X) < 0.1 {
		g.player.SetVelocity(Vec2f{0, g.player.GetVelocity().Y})
	}

	if math.Abs(g.player.GetVelocity().Y) < 0.1 {
		g.player.SetVelocity(Vec2f{g.player.GetVelocity().X, 0})
	}

	if g.player.GetVelocity().X > 0 {
		g.player.AddVelocity(Vec2f{-0.1, 0})
	}
	if g.player.GetVelocity().X < 0 {
		g.player.AddVelocity(Vec2f{0.1, 0})
	}
	if g.player.GetVelocity().Y > 0 {
		g.player.AddVelocity(Vec2f{0, -0.1})
	}
	if g.player.GetVelocity().Y < 0 {
		g.player.AddVelocity(Vec2f{0, 0.1})
	}

	if g.shoot_cooldown > 0 {
		g.shoot_cooldown--
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{46, 34, 47, 255})
	for _, o := range g.entity_list {
		o.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 160
}

// create collision slice from entity slice
func collisionFromEntity(eSlice []*em.Entity) []*cm.CollisionBox {
	var t []*cm.CollisionBox
	for _, e := range eSlice {
		t = append(t, e.GetCollider())
	}
	return t
}

func createBullet(g *Game) {
	mx, my := ebiten.CursorPosition()
	angle := math.Atan2(float64(my)-g.player.GetPosition().Y, float64(mx)-g.player.GetPosition().X)
	bs := image.Image{}
	bs.Init("assets/bullet.png", 16, 16)
	entity := em.Entity{}
	entity.Init(&bs, g.player.GetPosition(), Vec2f{5 * math.Cos(angle), 5 * math.Sin(angle)}, true)

	entity.SetRotation(angle)

	entity.MakeCollider(2, Vec2f{16, 16}, g.player.GetPosition(),
		func(id int) {
			if id == 1 {
				log.Printf("Killed enemy!\n")
			}
		},
	)

	g.entity_list = append(g.entity_list, &entity)
}

func main() {
	game := Game{}

	// create player entity

	player_sprite := am.AnimatedImage{}
	player_sprite.Init("assets/player.png", 500, 16, 16)
	go player_sprite.PlayLoop()
	player_sprite.SetPlaying(true)

	game.player.Init(&player_sprite, Vec2f{100.0, 100.0}, Vec2f{0, 0}, true)

	game.player.MakeCollider(0, Vec2f{16, 16}, Vec2f{100, 100},
		func(id int) {

		},
	)

	game.entity_list = append(game.entity_list, &game.player)

	// create enemy entity
	enemy := em.Entity{}
	enemy_sprite := am.AnimatedImage{}
	enemy_sprite.Init("assets/enemy.png", 500, 16, 16)
	go enemy_sprite.PlayLoop()
	enemy_sprite.SetPlaying(true)
	enemy.Init(&enemy_sprite, Vec2f{80, 80}, Vec2f{0, 0}, true)
	enemy.MakeCollider(1, Vec2f{16, 16}, Vec2f{80, 80},
		func(id int) {
			if id == 2 {
				// bullet

			}
		},
	)

	game.entity_list = append(game.entity_list, &enemy)

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
