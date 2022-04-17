package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type DebugUiSystem struct {
}

func (system *DebugUiSystem) Update() {
	rl.DrawFPS(10, 128)
}
