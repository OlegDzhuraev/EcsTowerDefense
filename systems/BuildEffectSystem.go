package systems

import (
	. "RtsGame/engine"
	. "RtsGame/engine/render"
	. "RtsGame/signals"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BuildEffectSystem struct {
}

func (system *BuildEffectSystem) Update() {
	if signal, ok := ecs.GetSignal[*SpawnFxSignal](); ok {
		ent := ecs.NewEntity()
		ent.Add(NewTransform(signal.Position))
		ent.Add(&TimedDestroy{Time: 1})
		ent.Add(&ModelRenderer{Color: rl.Green, Model: MakeCubeModel()})
		ent.Add(&CurveScale{Curve: FloatCurve{Keys: map[float32]float32{0: 2, 0.5: 0}}})
	}
}
