package entities

import (
	. "RtsGame/components"
	. "RtsGame/engine"
	. "RtsGame/engine/render"
	. "RtsGame/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewMineBuilding() ecs.Entity {
	e := ecs.NewEntity()

	e.Add(NewTransform(rl.Vector3{}))
	e.Add(&MineBuilding{MiningResource: "iron", AddPerSecond: 2, RequireField: true})
	e.Add(&ModelRenderer{Color: rl.Blue, Model: MakeCubeModel()})
	e.Add(&Damageable{MaxHealth: 200, Health: 200})
	e.Add(&Team{Id: 0})
	e.Add(&PlayerOwnedTag{})
	e.Add(&TargetTag{})

	return e
}
