package particle

import (
	. "bob/libs/types"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct{
	image *ebiten.Image
	lifetime int64 // unix mili
	position Vec2f
	velocity Vec2f
	dead bool
}

func (p *Particle) Update(){
	if time.Now().UnixMilli() >= p.lifetime {
		p.dead = true
	}
	p.position.AddEql(&p.velocity)
}

func (p *Particle) Draw(image *ebiten.Image){
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.position.X, p.position.Y)
	image.DrawImage(p.image, &opt)
}

type ParticleEmitter struct {
	image *ebiten.Image
	particles []*Particle
	position Vec2f
}

func CreateEmitter() ParticleEmitter{
	p := ParticleEmitter{}

	return p
}

func (p *ParticleEmitter) AddParticle(x_spread, y_spread float32, lifetime_spread int){
	t := Particle{}
	t.dead = false
	t.lifetime = time.Now().UnixMilli() + int64(rand.Intn(lifetime_spread))
	t.image = ebiten.NewImage(1, 1)
	t.image.Fill(color.RGBA{255, uint8(rand.Intn(255)), 0, 255})
	t.velocity.X = -float64(x_spread) + rand.Float64() * (float64(x_spread) - -float64(x_spread))
	t.velocity.Y = -float64(y_spread) + rand.Float64() * (float64(y_spread) - -float64(y_spread))
	t.position = p.position
	p.particles = append(p.particles, &t)
}

// TODO
// Fixa om två partiklar dör i samma cykel crash
func (p *ParticleEmitter) Update(){
	for i, par := range p.particles{
		if par.dead{
			if i >= len(p.particles){continue}
			p.particles[i] = p.particles[len(p.particles)-1]
			p.particles = p.particles[:len(p.particles)-1]
			continue
		}
		par.Update()
	}
}

func (p *ParticleEmitter) Draw(image *ebiten.Image){
	for _, par := range p.particles{
		par.Draw(image)
	}
}

func (p *ParticleEmitter) SetPosition(new_pos Vec2f){
	p.position = new_pos
}
