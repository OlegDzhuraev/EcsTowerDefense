package entities

import (
	. "RtsGame/components"
	. "RtsGame/engine"
	. "RtsGame/engine/render"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewTurret(parent ecs.Entity) ecs.Entity {
	e := ecs.NewEntity()
	tr := NewTransform(rl.Vector3{})
	tr.LocalPosition.Y = 1
	tr.Scale = rl.Vector3{X: 0.5, Y: 0.5, Z: 0.5}

	e.Add(tr)
	e.Add(&ChildEntity{Parent: parent})
	e.Add(&ModelRenderer{Color: rl.Yellow, Model: MakeCubeModel()})
	e.Add(&TargetSearcher{MaxDistance: 10})
	e.Add(&Team{Id: 0})
	e.Add(&Attack{Damage: 10, ReloadTime: 1, Distance: 8})

	return e
}
