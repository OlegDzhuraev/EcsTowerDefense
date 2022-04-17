package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Movable struct {
	Speed       float32
	Destination rl.Vector3
}
