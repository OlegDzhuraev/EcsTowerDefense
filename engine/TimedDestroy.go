package engine

import (
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TimedDestroy struct {
	Time float32
}

type TimedDestroySystem struct {
}

func (ts *TimedDestroySystem) Update() {
	filter := ecs.FilterWith[*TimedDestroy]()
	dTime := rl.GetFrameTime()

	for _, td := range filter {
		td.Time -= dTime

		if td.Time <= 0 {
			e := ecs.GetEntity(td)
			ecs.KillEntity(e)
		}
	}
}
