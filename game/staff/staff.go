package staff

import (
	"bob/game/bullet"
	"bob/libs/particle"
	. "bob/libs/types"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Staff struct {
	image    *ebiten.Image
	position Vec2f
	rotation float64
	radius   float64

	emitter particle.ParticleEmitter

	bullets []*bullet.Bullet
}

func CreateStaff(path string, radius float64) Staff {
	s := Staff{}
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal("Error loading in file to staff")
	}
	s.image = img
	s.radius = radius
	return s
}

func (s *Staff) Update(origin Vec2f) {
	mx, my := ebiten.CursorPosition()
	
	a := math.Atan2(float64(my) - origin.Y, float64(mx) - origin.X)

	s.rotation = a

	s.position.X = s.radius * math.Cos(a) + origin.X
	s.position.Y = s.radius * math.Sin(a) + origin.Y
	
	s.emitter.SetPosition(Vec2f{
		(s.radius + 3) * math.Cos(a) + origin.X,
		(s.radius + 3) * math.Sin(a) + origin.Y,
	})
	s.emitter.Update()
}


func (s *Staff) Draw(image *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(s.image.Bounds().Dx()/2), -float64(s.image.Bounds().Dy()/2))
	
	opt.GeoM.Rotate(s.rotation)
	
	opt.GeoM.Translate(s.position.X, s.position.Y)

	image.DrawImage(s.image, &opt)
	s.emitter.Draw(image)
}


func (s *Staff) Shoot(){
	s.emitter.AddParticle(1.5, 1.5, 200)
}
