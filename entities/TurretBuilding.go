package entities

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/engine/render"
	. "TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewTurretBuilding() ecs.Entity {
	e := ecs.NewEntity()

	e.Add(NewTransform(rl.Vector3{}))
	e.Add(&ModelRenderer{Color: rl.Red, Model: MakeCubeModel()})
	e.Add(&Damageable{MaxHealth: 100, Health: 100})
	e.Add(&TargetTag{})
	e.Add(&Team{Id: 0})
	e.Add(&PlayerOwnedTag{})

	NewTurret(e)

	return e
}
