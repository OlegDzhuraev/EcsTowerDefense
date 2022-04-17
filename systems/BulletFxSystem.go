package systems

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/tags"
	. "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BulletFxSystem struct {
	killDist float32
}

func (system *BulletFxSystem) Init() {
	system.killDist = 0.2
}

func (system *BulletFxSystem) Update() {
	_, transforms, movables := FilterWith3[*BulletFxTag, *Transform, *Movable]()

	for i, transform := range transforms {
		movable := movables[i]

		if rl.Vector3Distance(transform.Position, movable.Destination) < system.killDist {
			KillEntity(GetEntity(transform))
		}
	}
}
