package entities

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/engine/render"
	. "TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewHive(pos rl.Vector3) ecs.Entity {
	e := ecs.NewEntity()
	tr := NewTransform(pos)
	tr.Scale = rl.Vector3{X: 2, Y: 2, Z: 2}

	e.Add(tr)
	e.Add(&HiveTag{})
	e.Add(&ScoreForDestroy{Score: 1000})
	e.Add(&Damageable{MaxHealth: 300, Health: 300})
	e.Add(&ModelRenderer{Color: rl.Red, Model: MakeSphereModel()})
	e.Add(&Team{Id: 1})
	e.Add(&TargetTag{})

	return e
}
