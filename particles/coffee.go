package particles

import (
	"math"
	"math/rand"
	"time"
)

type Coffee struct {
	ParticleSystem
}

func reset(p *Particle, params *ParticleParams) {
	p.lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	p.speed = params.MaxSpeed * rand.Float64()

	maxX := math.Floor(float64(params.X) / 2)
	x := math.Max(-maxX, math.Min(rand.NormFloat64()*params.XScale, maxX))
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

func NewCoffee(width, height int, scale float64) Coffee {
	startTime := time.Now().UnixMilli()

	ascii := func(row, col int, counts [][]int) string {
		count := counts[row][col]
		if count < 2 {
			return " "
		}

		// if count == 1 {
		// 	if row > 0 && counts[row-1][col] > 0 {
		// 		return "."
		// 	}
		// 	return " "
		// }

		direction := row + int(((time.Now().UnixMilli()-startTime)/1000)%2)

		if direction%2 == 0 {
			return "}"
		}

		return "{"
	}

	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:       6000,
				MaxSpeed:      1.5,
				ParticleCount: 400,
				reset:         reset,
				ascii:         ascii,
				nextPosition:  nextPost,
				XScale:        scale,
				X:             width,
				Y:             height,
			},
		),
	}
}
