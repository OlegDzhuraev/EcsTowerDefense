package render

import (
	"RtsGame/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ModelRenderer struct {
	Color rl.Color
	Model rl.Model
}

func (mr *ModelRenderer) Render(t *engine.Transform) {
	// todo angle and rotation from transform ?
	// todo mesh instancing?
	rl.DrawModelEx(mr.Model, t.Position, rl.Vector3{Y: 1}, 0, t.Scale, mr.Color)
}

func (mr *ModelRenderer) GetOrder() int16      { return 0 }
func (mr *ModelRenderer) SetOrder(order int16) {}

func (mr *ModelRenderer) GetColor() rl.Color      { return mr.Color }
func (mr *ModelRenderer) SetColor(color rl.Color) { mr.Color = color }
