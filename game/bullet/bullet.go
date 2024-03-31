package bullet

import (
	cm "bob/libs/collision"
	"bob/libs/particle"
	. "bob/libs/types"
	"math/rand"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	position  Vec2f
	velocity  Vec2f
	emitter   particle.ParticleEmitter
	collider  cm.CollisionBox
	life_time int64
	Is_dead   bool
}

func CreateBullet(position Vec2f, angle float64, life_time int64) Bullet {
	b := Bullet{}
	b.Is_dead = false
	b.position = position
	b.life_time = life_time
	b.velocity.Y = 2 * math.Sin(angle)
	b.velocity.X = 2 * math.Cos(angle)
	b.collider = cm.CreateCB(2, Vec2f{1, 1}, position,
		func(id int) {
			log.Printf("Bullet hit %v\n", id)
		},
	)
	return b
}

func (b *Bullet) Draw(image *ebiten.Image) {
	if b.Is_dead {
		return
	}
	b.emitter.Draw(image)
}

func (b *Bullet) Update() {
	if b.Is_dead {
		return
	}
	if time.Now().UnixMilli() >= b.life_time {
		b.Is_dead = true
		return
	}
	b.position.AddEql(&b.velocity)
	b.collider.SetPosition(b.position)
	b.emitter.SetPosition(b.position)
	b.emitter.AddParticle(1, 1, 100, color.RGBA{0, uint8(rand.Intn(255)), 255, 255})
	b.emitter.Update()
}
