package render

import (
	. "RtsGame/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// todo finish

// ColorTween is unfinished now
type ColorTween struct {
	Target        Renderer
	TargetColor   rl.Color
	Duration      float32
	Callback      func()
	startDuration float32
	startColor    rl.Color
}

func (ct *ColorTween) Update() {
	if ct.Duration > ct.startDuration {
		ct.startDuration = ct.Duration
		ct.startColor = ct.Target.GetColor()
	}

	ct.Duration -= rl.GetFrameTime()
	col := Lerp(ct.startColor, ct.TargetColor, 1-(ct.Duration/ct.startDuration))
	ct.Target.SetColor(col)

	if ct.Duration <= 0 {
		ct.Target.SetColor(ct.TargetColor)
		// todo KillBehavior(ct)

		if ct.Callback != nil {
			ct.Callback()
		}
	}
}
