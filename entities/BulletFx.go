package entities

import (
	. "RtsGame/components"
	. "RtsGame/engine"
	. "RtsGame/engine/render"
	. "RtsGame/tags"
	. "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBulletEffect(pos rl.Vector3, target rl.Vector3) Entity {
	e := NewEntity()
	tr := NewTransform(pos)
	tr.Scale = rl.Vector3{X: 0.2, Y: 0.2, Z: 0.2}

	e.Add(tr)
	e.Add(&BulletFxTag{})
	e.Add(&Movable{Speed: 25, Destination: target})
	e.Add(&ModelRenderer{Color: rl.White, Model: MakeSphereModel()})

	return e
}
