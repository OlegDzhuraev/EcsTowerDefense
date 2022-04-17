package systems

import (
	. "TowerDefenseTalosEcs/entities"
	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type GenerateResourcesSystem struct {
}

func (system *GenerateResourcesSystem) Init() {
	mapSize := 100.
	noiseScale := 0.3

	p := perlin.NewPerlin(2, 2, 3, rand.Int63())

	for x := -mapSize / 2; x < mapSize/2; x += 1 {
		for y := -mapSize / 2; y < mapSize/2; y += 1 {
			val := p.Noise2D(x/noiseScale, y/noiseScale)

			if val > 0.6 {
				NewResourceField(rl.Vector3{X: float32(x), Z: float32(y)})
			}
		}
	}
}
