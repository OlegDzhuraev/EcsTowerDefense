package systems

import (
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/engine/render"
	. "TowerDefenseTalosEcs/oneframes"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShootEffectSystem struct {
}

func (system *ShootEffectSystem) Update() {
	shoots := ecs.FilterWith[*ShootOneFrame]()

	for _, signal := range shoots {
		ent := ecs.NewEntity()
		ent.Add(NewTransform(signal.Position))
		ent.Add(&TimedDestroy{Time: 0.5})
		// todo custom renderer for effects like this, directly via rl.Draw to optimize
		ent.Add(&ModelRenderer{Color: rl.Yellow, Model: MakeSphereModel()})
		ent.Add(&CurveScale{Curve: FloatCurve{Keys: map[float32]float32{0: 0.8, 0.5: 0}}})
	}
}
