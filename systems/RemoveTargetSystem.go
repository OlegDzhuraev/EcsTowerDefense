package systems

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RemoveTargetSystem struct {
}

func (system *RemoveTargetSystem) Update() {
	targetSearchers, transforms := FilterWith2[*TargetSearcher, *Transform]()

	for i, searcher := range targetSearchers {
		if !IsAlive(searcher.Target) {
			continue
		}

		transform := transforms[i]

		if targetTransform, ok := GetComponent[*Transform](searcher.Target); ok {
			dist := rl.Vector3Distance(transform.Position, targetTransform.Position)

			if dist > searcher.MaxDistance {
				searcher.Target = 0
			}
		}
	}
}
