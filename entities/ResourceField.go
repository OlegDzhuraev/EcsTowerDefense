package entities

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/engine/render"
	. "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

const maxResources = 360
const minResources = 90

const minResScale = 0.45
const maxResScale = 1.8

func NewResourceField(pos rl.Vector3) Entity {
	e := NewEntity()

	rndResourcesPercent := rand.Float32()
	resources := int32(LerpFloat(minResources, maxResources, rndResourcesPercent))
	scale := GetResourceFieldScale(resources)

	tr := NewTransform(rl.Vector3Subtract(pos, rl.Vector3{Y: 0.49})) // todo it can affect mine search distance, view should be drawn separated
	tr.Scale = rl.Vector3{X: scale, Y: 0.1, Z: scale}

	e.Add(tr)
	e.Add(&ResourceField{ResourceType: "iron", Amount: resources})
	e.Add(&ModelRenderer{Color: rl.Color{R: 12, G: 28, B: 52, A: 255}, Model: MakeSphereModel()})

	return e
}

func GetResourceFieldScale(resourcesAmount int32) float32 {
	percents := InverseLerpFloat(minResources, maxResources, float32(resourcesAmount))
	return LerpFloat(minResScale, maxResScale, percents)
}
