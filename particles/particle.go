
package particles

import (
	"math"
	"strings"
	"time"
)

type Particle struct {
	lifetime int64
	speed    float64
	X        float64
	Y        float64
}

type ParticleParams struct {
	MaxLife       int64
	MaxSpeed      float64
	ParticleCount int
	X             int
	Y             int
	nextPosition  NextPosition
	ascii         Ascii
	reset         Reset
}

type NextPosition func(particle *Particle, deltaMS int64)
type Ascii func(row, col int, counts [][]int) string
type Reset func(particle *Particle, params *ParticleParams)

type ParticleSystem struct {
	ParticleParams
	particles []*Particle
	lastTime  int64
}

func NewParticleSystem(params ParticleParams) ParticleSystem {
  particles := make([]* Particle, 0)
  for i:=0; i<params.ParticleCount; i++{
    particles = append(particles, &Particle{})
  }
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
    particles: particles,
	}
}

func (ps *ParticleSystem) Start() {
	for _, p := range ps.particles {
		ps.reset(p, &ps.ParticleParams)
	}
}

func (ps *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	delta := now - ps.lastTime
	ps.lastTime = now

	for _, p := range ps.particles {
		ps.nextPosition(p, delta)

		if p.Y >= float64(ps.Y) || p.X >= float64(ps.X) || p.lifetime <= 0{
			ps.reset(p, &ps.ParticleParams)
		}
	}
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


func (ps *ParticleSystem) Display() string {
	counts := make([][]int, ps.Y)

	for row := range counts {
		counts[row] = make([]int, ps.X)
	}

	for _, p := range ps.particles {
		row := int(math.Floor(p.Y))
		col := int(math.Floor(p.X))
		if row >= 0 && row < len(counts) && col >= 0 && col < len(counts[0]) {
			counts[row][col]++
		}
	}

	out := make([][]string, len(counts))

	for r := range counts {
		out[r] = make([]string, len(counts[r]))
		for c := range counts[r] {
			out[r][c] = ps.ascii(r, c, counts)
		}
	}

	Reverse(out)

	outStr := make([]string, len(out))
	for i, row := range out {
		outStr[i] = strings.Join(row, "")
	}

	return strings.Join(outStr, "\n")
}


