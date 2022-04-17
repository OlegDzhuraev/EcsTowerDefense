package systems

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyMoveSystem struct {
}

func (system *EnemyMoveSystem) Update() {
	movables, tsearchers := ecs.FilterW2Excl1[*Movable, *TargetSearcher, *PlayerOwnedTag]()

	for i, m := range movables {
		movableEnt := ecs.GetEntity(m)
		target := tsearchers[i].Target

		if targetTransform, ok := ecs.GetComponent[*Transform](target); ok {
			if movableTransform, ok2 := ecs.GetComponent[*Transform](movableEnt); ok2 {
				offset := rl.Vector3Scale(GetDirection3d(movableTransform.Position, targetTransform.Position), 1)
				m.Destination = rl.Vector3Subtract(targetTransform.Position, offset)
			}
		}
	}
}
