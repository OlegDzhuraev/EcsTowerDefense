package render

import (
	. "RtsGame/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer interface {
	Render(t *Transform)
	SetColor(color rl.Color)
	GetColor() rl.Color
}
