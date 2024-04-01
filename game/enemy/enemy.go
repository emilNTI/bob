package enemy

import (
	"bob/game/player"
	cm "bob/libs/collision"
	"bob/libs/drawble"
	. "bob/libs/types"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SKELETON_SPEED = 0.5
	SKELETON_MAX_SPEED = 0.5
)

type Enemy interface {
	Update()
	Draw(image *ebiten.Image)

	Attack()
}

// Skeleton
type Skeleton struct {
	drawble drawble.Drawble
	position Vec2f
	velocity Vec2f

	// collision
	collider cm.CollisionBox

	flip_x bool

	target *player.Player
}

func CreateSkeleton(drawble drawble.Drawble, position Vec2f, target *player.Player) Skeleton {
	s := Skeleton{}
	s.drawble = drawble
	s.position = position
	s.velocity = Vec2f{0, 0}
	s.target = target
	s.collider = cm.CreateCB(3, Vec2f{16, 16}, Vec2f{0, 0},
		func(id int) {
			log.Printf("Skeletion collide with %v\n", id)
		},
	)

	return s
}


// TODO
// FIX BUGGY MOVMENT
func (s *Skeleton) Update() {
	pc := s.target.GetCenter()
	sc := s.GetCenter()
	if sc.X > pc.X && s.velocity.X > -SKELETON_MAX_SPEED {
		s.velocity.X += -SKELETON_SPEED
	}
	if sc.X < pc.X && s.velocity.X < SKELETON_MAX_SPEED{
		s.velocity.X += SKELETON_SPEED
	}
	if sc.Y > pc.Y && s.velocity.Y > -SKELETON_MAX_SPEED{
		s.velocity.Y += -SKELETON_SPEED
	}
	if sc.Y < pc.Y && s.velocity.Y < SKELETON_MAX_SPEED{
		s.velocity.Y += SKELETON_SPEED
	}

	// set flip x
	s.flip_x = math.Signbit(s.velocity.X)
	
	s.position.AddEql(&s.velocity)

	s.collider.SetPosition(s.position)
}

func (s *Skeleton) Draw(image *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}

	if s.flip_x {
		opt.GeoM.Scale(-1, 1)
		opt.GeoM.Translate(s.drawble.GetSize().X, 0)
	}

	opt.GeoM.Translate(s.position.X, s.position.Y)

	image.DrawImage(s.drawble.GetImage(), &opt)
}

func (s *Skeleton) Attack(player *player.Player) {}

func (s *Skeleton) GetCenter() Vec2f {
	return Vec2f{
		s.position.X + s.drawble.GetSize().X / 2,
		s.position.Y + s.drawble.GetSize().Y / 2,
	}
}

func (s *Skeleton) GetHitBox() *cm.CollisionBox {
	return &s.collider
}
