package entities

import (
	. "RtsGame/components"
	. "RtsGame/engine"
	. "RtsGame/engine/render"
	. "RtsGame/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewEnemy() ecs.Entity {
	e := ecs.NewEntity()
	tr := NewTransform(rl.Vector3{})
	tr.Scale = rl.Vector3{X: 0.5, Y: 0.5, Z: 0.5}

	e.Add(tr)
	e.Add(&Movable{Speed: 1.5})
	e.Add(&ModelRenderer{Color: rl.Violet, Model: MakeSphereModel()})
	e.Add(&ScoreForDestroy{Score: 100})
	e.Add(&Damageable{MaxHealth: 100, Health: 100})
	e.Add(&Attack{Damage: 10, Distance: 1.25, ReloadTime: 1})
	e.Add(&Team{Id: 1})
	e.Add(&TargetTag{})
	e.Add(&TargetSearcher{MaxDistance: 10})

	return e
}
