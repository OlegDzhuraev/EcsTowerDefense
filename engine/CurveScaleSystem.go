package engine

import (
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CurveScale struct {
	Curve    FloatCurve
	Progress float32
}

type CurveScaleSystem struct {
}

func (system *CurveScaleSystem) Update() {
	curves, transforms := ecs.FilterWith2[*CurveScale, *Transform]()

	dTime := rl.GetFrameTime()

	for i, c := range curves {
		t := transforms[i]

		c.Progress += dTime
		t.Scale = rl.Vector3Scale(rl.Vector3{X: 1, Y: 1, Z: 1}, c.Curve.Evaluate(c.Progress))

		if c.Progress >= c.Curve.GetTimeLength() {
			ecs.DelComponent[*CurveScale](ecs.GetEntity(c))
		}
	}
}
