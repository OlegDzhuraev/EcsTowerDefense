package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Transform struct {
	Position      rl.Vector3
	LocalPosition rl.Vector3
	Rotation      rl.Vector3
	Scale         rl.Vector3
}

func (t *Transform) SetPosition(pos rl.Vector3) { t.Position = pos }
func (t *Transform) GetTransform() *Transform   { return t }

func NewTransform(pos rl.Vector3) *Transform {
	return &Transform{
		Position: pos,
		Scale:    rl.Vector3{X: 1, Y: 1, Z: 1},
	}
}
