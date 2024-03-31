package bullet

import (
	"bob/libs/particle"
	. "bob/libs/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	position Vec2f
	velocity Vec2f
	emitter particle.ParticleEmitter
}

func CreateBullet(position, velocity Vec2f) Bullet {
	b := Bullet{}
	b.position = position
	b.velocity = velocity

	return b
}

func (b *Bullet) Draw(image *ebiten.Image) {
	b.emitter.Draw(image)
}

func (b *Bullet) Update() {
	b.position.AddEql(&b.velocity)
	b.emitter.SetPosition(b.position)
	b.emitter.AddParticle(1, 1, 100)
	b.Update()
}
