package entities

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/engine/render"
	. "TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewHeadquartersBuilding() ecs.Entity {
	e := ecs.NewEntity()
	tr := NewTransform(rl.Vector3{})
	tr.Scale = rl.Vector3{X: 2, Y: 1, Z: 2}

	e.Add(tr)
	e.Add(&MineBuilding{MiningResource: "iron", AddPerSecond: 2, RequireField: false})
	e.Add(&ModelRenderer{Color: rl.Green, Model: MakeCubeModel()})
	e.Add(&Damageable{MaxHealth: 250, Health: 250})
	e.Add(&Team{Id: 0})
	e.Add(&TargetTag{})
	e.Add(&PlayerOwnedTag{})
	e.Add(&HeadquartersTag{})

	e2 := ecs.NewEntity()
	tr2 := NewTransform(rl.Vector3{})
	tr2.LocalPosition.Y = 0.75
	tr2.Scale = rl.Vector3{X: 1, Y: 0.5, Z: 1}

	e2.Add(tr2)
	e2.Add(&ChildEntity{Parent: e})
	e2.Add(&ModelRenderer{Color: rl.Blue, Model: MakeCubeModel()})

	return e
}
