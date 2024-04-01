package player

import (
	"bob/game/staff"
	cm "bob/libs/collision"
	d "bob/libs/drawble"
	. "bob/libs/types"
	"math"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MAX_WALK_SPEED = 1.5
	WALK_SPEED     = 0.3
	BREAK_SPEED    = 0.1
	SNAP_TO_ZERO   = 0.1
)

type Player struct {
	drawble d.Drawble // sprite
	// collision
	env_collider cm.CollisionBox
	sec_collider cm.CollisionBox
	// info
	position Vec2f
	velocity Vec2f
	flip_x bool

	// other
	wizard_staff staff.Staff
}

func CreatePlayer(drawble d.Drawble, env_c, sec_c cm.CollisionBox, pos Vec2f) Player {
	p := Player{}
	p.drawble = drawble
	p.env_collider = env_c
	p.sec_collider = sec_c
	p.position = pos

	// create staff
	p.wizard_staff = staff.CreateStaff("assets/staff.png", 10)
	return p
}

func (p *Player) Update() {
	// movment
	if ebiten.IsKeyPressed(ebiten.KeyA) && p.velocity.X > -MAX_WALK_SPEED {
		p.velocity.AddEql(&Vec2f{-WALK_SPEED, 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && p.velocity.X < MAX_WALK_SPEED {
		p.velocity.AddEql(&Vec2f{WALK_SPEED, 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.velocity.Y > -MAX_WALK_SPEED {
		p.velocity.AddEql(&Vec2f{0, -WALK_SPEED})
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && p.velocity.Y < MAX_WALK_SPEED {
		p.velocity.AddEql(&Vec2f{0, WALK_SPEED})
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.wizard_staff.Shoot()
	}
	
	// break player
	if p.velocity.X > 0.0 {
		p.velocity.X += -BREAK_SPEED
	}
	if p.velocity.X < 0.0 {
		p.velocity.X += BREAK_SPEED
	}
	if p.velocity.Y > 0.0 {
		p.velocity.Y += -BREAK_SPEED
	}
	if p.velocity.Y < 0.0 {
		p.velocity.Y += BREAK_SPEED
	}

	// snap to zero if close enough
	if math.Abs(p.velocity.X) <= SNAP_TO_ZERO{
		p.velocity.X = 0
	}
	if math.Abs(p.velocity.Y) <= SNAP_TO_ZERO{
		p.velocity.Y = 0
	}
	
	// update position
	p.position.AddEql(&p.velocity)

	// update colliders position
	p.env_collider.SetPosition(Vec2f{p.position.X + 4, p.position.Y + 13})
	p.sec_collider.SetPosition(Vec2f{p.position.X + 4, p.position.Y})

	// check if should flip player
	mx, _ := ebiten.CursorPosition()
	if mx > int(math.Round(p.drawble.GetSize().X / 2 + p.position.X)) {
		p.flip_x = false
	} else {p.flip_x = true}

	// update staff
	p.wizard_staff.Update(p.GetCenter())
}

func (p *Player) FlipX(should_flip bool) {
	p.flip_x = should_flip
}

func (p *Player) Draw(image *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	if p.flip_x {
		opt.GeoM.Scale(-1, 1)
		opt.GeoM.Translate(16, 0)
	}
	
	opt.GeoM.Translate(p.position.X, p.position.Y)
	//p.env_collider.Draw(image, color.RGBA{255, 0, 0, 255})
	//p.sec_collider.Draw(image, color.RGBA{0, 255, 0, 50})
	image.DrawImage(p.drawble.GetImage(), &opt)

	p.wizard_staff.Draw(image)
}

func (p *Player) GetPosition() Vec2f{
	return p.position
}

func (p *Player) GetCenter() Vec2f{
	return Vec2f{
		p.position.X + p.drawble.GetSize().X / 2,
		p.position.Y + p.drawble.GetSize().Y / 2,
	}
}

func (p *Player) GetHitBox() *cm.CollisionBox {
	return &p.sec_collider
}
