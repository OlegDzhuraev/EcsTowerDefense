package systems

import (
	. "RtsGame/components"
	"RtsGame/engine"
	"RtsGame/engine/render"
	"RtsGame/entities"
	"RtsGame/settings"
	. "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type MineSystem struct {
	maxDistToField float32
}

func (system *MineSystem) Init() {
	system.maxDistToField = 1.25
}

func (system *MineSystem) Update() {
	mines := FilterWith[*MineBuilding]()
	dTime := rl.GetFrameTime()

	for _, mine := range mines {
		mine.Timer -= dTime

		if mine.Timer <= 0 {
			addValue := mine.AddPerSecond
			mine.Timer = 1

			if mine.RequireField {
				addValue = int32(math.Min(float64(mine.Field.Amount), float64(mine.AddPerSecond)))
				mine.Field.Amount -= addValue

				if fieldTr, ok := GetComponent[*engine.Transform](GetEntity(mine.Field)); ok {
					scale := entities.GetResourceFieldScale(mine.Field.Amount)
					fieldTr.Scale = rl.Vector3{X: scale, Y: fieldTr.Scale.Y, Z: scale}
				}

				if mine.Field.Amount <= 0 {
					e := GetEntity(mine)
					if mr, ok := GetComponent[*render.ModelRenderer](e); ok {
						mr.SetColor(rl.Gray)
					}
					DelComponent[*MineBuilding](e)
				}
			}

			settings.PlayerResources[mine.MiningResource] += addValue
		}
	}
}
