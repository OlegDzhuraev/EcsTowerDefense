package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type InputKeyAxis struct {
	KeyA int32
	KeyB int32
}

var HorizontalAxis = InputKeyAxis{rl.KeyA, rl.KeyD}
var VerticalAxis = InputKeyAxis{rl.KeyS, rl.KeyW}

func (axis InputKeyAxis) GetValue() float32 {
	if rl.IsKeyDown(axis.KeyA) {
		return -1
	} else if rl.IsKeyDown(axis.KeyB) {
		return 1
	}

	return 0
}
