package particles

import (
	"math"
	"math/rand"

)
type Coffee struct {
  ParticleSystem
}
func ascii(row, col int, counts [][]int) string {
	count := counts[row][col]
	if count < 3 {
		return " "
	}
	if count < 6 {
		return "."
	}
	if count < 9 {
		return ":"
	}
	if count < 12 {
		return "{"
	}
	return "}"
}

func reset(p *Particle, params *ParticleParams) {
	p.lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	p.speed = params.MaxSpeed * rand.Float64()

	maxX := math.Floor(float64(params.X) / 2)
	x := math.Max(-maxX, math.Min(rand.NormFloat64(), maxX))
	p.X = x + maxX
	p.Y = 0
}

func nextPost(p *Particle, deltaMS int64) {
	p.lifetime -= deltaMS
	if p.lifetime <= 0 {
		return
	}
  percent := (float64(deltaMS) / 1000.0)
	p.Y += p.speed * percent
}

func NewCoffee(width, height int) Coffee {


	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:       7000,
				MaxSpeed:      0.5,
				ParticleCount: 100,
				reset:         reset,
				ascii:         ascii,
				nextPosition:  nextPost,
				X:             width,
				Y:             height,
			},
		),
	}
}

