package systems

import (
	. "RtsGame/components"
	"RtsGame/engine"
	. "RtsGame/tags"
	. "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyMoveSystem struct {
}

func (system *EnemyMoveSystem) Update() {
	movables, tsearchers := FilterW2Excl1[*Movable, *TargetSearcher, *PlayerOwnedTag]()

	for i, m := range movables {
		movableEnt := GetEntity(m)
		target := tsearchers[i].Target

		if targetTransform, ok := GetComponent[*engine.Transform](target); ok {
			if movableTransform, ok2 := GetComponent[*engine.Transform](movableEnt); ok2 {
				offset := rl.Vector3Scale(engine.GetDirection3d(movableTransform.Position, targetTransform.Position), 1)
				m.Destination = rl.Vector3Subtract(targetTransform.Position, offset)
			}
		}
	}
}
