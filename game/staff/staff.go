package staff

import (
	"bob/game/bullet"
	"bob/libs/particle"
	. "bob/libs/types"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

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

	cool_down int64
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

	a := math.Atan2(float64(my)-origin.Y, float64(mx)-origin.X)

	s.rotation = a

	s.position.X = s.radius*math.Cos(a) + origin.X
	s.position.Y = s.radius*math.Sin(a) + origin.Y

	s.emitter.SetPosition(Vec2f{
		(s.radius+3)*math.Cos(a) + origin.X,
		(s.radius+3)*math.Sin(a) + origin.Y,
	})
	s.emitter.Update()

	for i, b := range s.bullets {
		if b.Is_dead {
			if i >= len(s.bullets) {
				continue
			}
			s.bullets[i] = s.bullets[len(s.bullets)-1]
			s.bullets = s.bullets[:len(s.bullets)-1]
		}
		b.Update()
	}
}

func (s *Staff) Draw(image *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(s.image.Bounds().Dx()/2), -float64(s.image.Bounds().Dy()/2))

	opt.GeoM.Rotate(s.rotation)

	opt.GeoM.Translate(s.position.X, s.position.Y)

	image.DrawImage(s.image, &opt)
	s.emitter.Draw(image)

	for _, b := range s.bullets {
		b.Draw(image)
	}
}

func (s *Staff) Shoot() {
	// check cool_down
	if time.Now().UnixMilli() < s.cool_down {
		return
	}
	// set cool_down
	s.cool_down = time.Now().UnixMilli() + 50

	// create bullet
	b := bullet.CreateBullet(s.emitter.GetPosition(), s.rotation, time.Now().UnixMilli()+2000)
	s.bullets = append(s.bullets, &b)

	// create particles on staff
	for i := 0; i < 10; i++ {
		s.emitter.AddParticle(.5, .5, 200, color.RGBA{255, uint8(rand.Intn(255)), 0, 255})
	}
}
