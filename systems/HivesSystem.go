package systems

import (
	"RtsGame/engine"
	"RtsGame/entities"
	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type HivesSystem struct {
}

func (system *HivesSystem) Init() {
	mapSize := 70.
	noiseScale := float32(0.3)
	var maxHives int32 = 10

	p := perlin.NewPerlin(2, 2, 3, int64(rand.Int()))

	var spawned int32 = 0
	for spawned < maxHives {
		x := engine.LerpFloat(float32(-mapSize/2), float32(mapSize/2), rand.Float32())
		y := engine.LerpFloat(float32(-mapSize/2), float32(mapSize/2), rand.Float32())
		var centerDist = rl.Vector2Distance(rl.Vector2{X: x, Y: y}, rl.Vector2{})
		if centerDist < 15 {
			continue // todo make it better, there can be a lot of useless iterations
		}

		val := p.Noise2D(float64(x/noiseScale), float64(y/noiseScale)) + 0.5

		if val > 0.985 {
			entities.NewHive(rl.Vector3{X: x, Z: y})
			spawned++
		}
	}
}
