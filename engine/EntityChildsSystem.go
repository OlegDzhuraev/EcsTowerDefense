package engine

import (
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ChildrenSystem struct {
}

func (system *ChildrenSystem) Update() {
	AliveCheck()
	TransformUpdate()
}

func AliveCheck() {
	children := ecs.FilterWith[*ChildEntity]()

	for _, ch := range children {
		if !ecs.IsAlive(ch.Parent) {
			ecs.KillEntity(ecs.GetEntity(ch))
		}
	}
}

func TransformUpdate() {
	children, transforms := ecs.FilterWith2[*ChildEntity, *Transform]()

	for i, child := range children {
		if parentTr, ok := ecs.GetComponent[*Transform](child.Parent); ok {
			transforms[i].Position = rl.Vector3Add(parentTr.Position, transforms[i].LocalPosition)
		}
	}
}
