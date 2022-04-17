package systems

import (
	. "RtsGame/components"
	. "RtsGame/engine"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MoveSystem struct {
	stopDistance float32
}

func (system *MoveSystem) Init() {
	system.stopDistance = 0.15
}

func (system *MoveSystem) Update() {
	movables, transforms := ecs.FilterWith2[*Movable, *Transform]()
	dTime := rl.GetFrameTime()

	for i, movable := range movables {
		transform := transforms[i]

		if rl.Vector3Distance(transform.Position, movable.Destination) > system.stopDistance {
			dir := GetDirection3d(transform.Position, movable.Destination)
			dir = rl.Vector3Normalize(dir)
			scaled := rl.Vector3Scale(dir, dTime*movable.Speed)
			transform.Position = rl.Vector3Add(transform.Position, scaled)
		}
	}
}
